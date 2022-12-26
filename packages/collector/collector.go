package Collector

import (
	"bytes"
	"io"
	"log"
	"net"
	"time"
	"strconv"
	"strings"
	"encoding/csv"
	"encoding/xml"

	"chicha/packages/config"
	"chicha/packages/data"
	"chicha/packages/proxy"
)
//initial function
func init() {
	//set microsecond resolution for logging:
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}

func IsValidXML(data []byte) bool {
	return xml.Unmarshal(data, new(interface{})) == nil
}


func processConnection(connection net.Conn) {
	defer connection.Close()
	var tempDelay time.Duration // how long to sleep on accept failure
	var rawData Data.RawData

	// Read connection in lap
	for {
		buf := make([]byte, 1024)
		size, err := connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				//log.Println("conn.Read(buf) error:", err)
				//log.Println("Message EOF detected - closing LAN connection.")
				break
			}

			networkError, ok := err.(*net.OpError) 
			if ok && networkError.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay = 2 * tempDelay
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				log.Printf("Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}

			break
		}
		tempDelay = 0

		data := buf[:size]

		remoteIPAddress := connection.RemoteAddr().(*net.TCPAddr).IP.String()

		//various data formats processing (text csv, xml) start:
		if !IsValidXML(data) {
			// CSV data processing
			csvData, err := csv.NewReader(bytes.NewReader(data)).ReadAll()
			if err != nil {
				log.Println(err)
				return
			}

			for _, csvField := range csvData {
				if len(csvField)==3 || len(csvField)==4 {
					// Prepare antenna position
					antennaPosition, err := strconv.ParseInt(string(strings.TrimSpace(csvField[2])), 10, 64)
					if err != nil {
						log.Println("Recived incorrect Antenna position CSV value:", err)
						continue
					}
					rawData.DiscoveryUnixTime, err = strconv.ParseInt(string(strings.TrimSpace(csvField[1])), 10, 64)
					if err != nil {
						log.Println("Recived incorrect discovery unix time CSV value:", err)
						continue
					}
					rawData.TagID = string(strings.TrimSpace(csvField[0]))
					rawData.Antenna = uint8(antennaPosition)
					if len(csvField) == 3 {
						//reader connection without proxy
						rawData.ReaderIP = remoteIPAddress

						//Debug all received data from RFID reader
						log.Printf("TAG=%s, TIME=%d, Reader-IP=%s, Reader-ANT=%d\n", rawData.TagID, rawData.DiscoveryUnixTime, rawData.ReaderIP, rawData.Antenna)

					} else if len(csvField) == 4 {
						//proxy connection
						if net.ParseIP(string(strings.TrimSpace(csvField[3]))) != nil {
							//sending data with proxy
							rawData.ReaderIP = string(strings.TrimSpace(csvField[3]))
							rawData.ProxyIP = remoteIPAddress
						} else {
							//sending csvField[3] is not an IP address
							rawData.ReaderIP = remoteIPAddress
						}
						//Debug all received data from PROXY
						log.Printf("TAG=%s, TIME=%d, Reader-IP=%s, Reader-Antenna=%d, Proxy-IP=%s\n", rawData.TagID, rawData.DiscoveryUnixTime, rawData.ReaderIP, rawData.Antenna, rawData.ProxyIP)

					}
					//create a proxy task if needed:
					if Config.PROXY_ADDRESS != "" {
						Proxy.ProxyTask <- rawData
					}
				}
			}
		} else {
			// XML data processing
			err := xml.Unmarshal(data, &rawData)
			if err != nil {
				log.Println("xml.Unmarshal ERROR:", err)
				continue
			}
			loc, err := time.LoadLocation(Config.TIME_ZONE)
			if err != nil {
				log.Println(err)
				continue
			}
			xmlTimeFormat := `2006/01/02 15:04:05.000`
			discoveryTime, err := time.ParseInLocation(xmlTimeFormat, string(rawData.DiscoveryUnixTime), loc)
			if err != nil {
				log.Println("time.ParseInLocation ERROR:", err)
				continue
			}
			rawData.DiscoveryUnixTime = discoveryTime.UnixNano()/int64(time.Millisecond)
			rawData.TagID = strings.ReplaceAll(rawData.TagID, " ", "")
			rawData.Antenna = uint8(rawData.Antenna)

			if net.ParseIP(rawData.ReaderIP) != nil {
				//connection received from proxy (not from reader).
				rawData.ProxyIP = remoteIPAddress

				//Debug all received data from PROXY
				log.Printf("TAG=%s, TIME=%d, Reader-IP=%s, Reader-Antenna=%d, Proxy-IP=%s\n", rawData.TagID, rawData.DiscoveryUnixTime, rawData.ReaderIP, rawData.Antenna, rawData.ProxyIP)
			} else {
				//connection received from reader (not from proxy)
				//Debug all received data from RFID reader
				log.Printf("TAG=%s, TIME=%d, Reader-IP=%s, Reader-ANT=%d\n", rawData.TagID, rawData.DiscoveryUnixTime, rawData.ReaderIP, rawData.Antenna)
			}
			//create a proxy task if needed (via Proxy.ProxyTask channel):
			if Config.PROXY_ADDRESS != "" {
				//send rawData to Proxy.ProxyTask channel
				Proxy.ProxyTask <- rawData
			}
		}

		/*
		if len(laps) == 0 {
			//laps buffer empty - recreate last race from db:
			log.Println("laps buffer empty - recreate last race from db")
			laps, err = Database.GetCurrentRaceDataFromDB()
			if err == nil {
				log.Printf("laps buffer recreated with %d records from db.\n", len(laps))
				go Database.addNewLapToLapsBuffer(rawData)
			} else {
				log.Println("laps buffer recreation failed with:", err)
				go Database.addNewLapToLapsBuffer(rawData)
			}
		}

		if len(laps) > 0 {
			// Add current Lap to Laps buffer
			go Database.addNewLapToLapsBuffer(rawData)
		}

		*/
	}
}

// Start data collector from RFID readers.
func StartDataCollector() {

	// Start listener
	collector, err := net.Listen("tcp", Config.COLLECTOR_LISTENER_ADDRESS)
	if err != nil {
		log.Panicln("Error: collector can't start. ", err)
	}
	defer collector.Close()


	// Listen new connections
	for {
		connection, err := collector.Accept()
		if err != nil {
			if err != io.EOF {
				log.Panicln("tcp connection error:", err)
			}
		}

		go processConnection(connection)
	}
}

