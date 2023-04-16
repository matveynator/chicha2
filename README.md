<p align="left">
    <img property="og:image" src="https://repository-images.githubusercontent.com/577755312/57f67b11-437b-448f-b53e-cf47165612c2" width="25%">
</p>

# Chicha - the competition timekeeper (chronograph). Version 2.0

Free chronograf for runners, bycicles, motorcycles, carts, cars, trucks, atv and other types of race competitions. 
UHF-RFID compatible.

<p align="left">
    <img property="og:image" src="https://repository-images.githubusercontent.com/368199185/e26c553e-b23e-4bae-b4d2-c2df502e9f04" width="75%">
</p>


## [Demo: http://chicha.zabiyaka.net](http://chicha.zabiyaka.net/)


## Конфигурационные опции (необязательные):
```
chicha -h
Usage of chicha:
  -average
    	Calculate average results instead of minimal results.
  -average-duration duration
    	Duration to calculate average results. Results passed to reader during this duration will be calculated as average result. (default 1s)
  -collector string
    	Provide IP address and port to collect and parse data from RFID and timing readers. (default "0.0.0.0:4000")
  -db-path string
    	Provide path to writable directory to store database data. (default ".")
  -db-save-interval duration
    	Duration to save data from memory to database (disk). Setting duration too low may cause unpredictable performance results. (default 30s)
  -db-type string
    	Select db type: sqlite / genji / postgres (default "genji")
  -lap-time duration
    	Minimal lap time duration. Results smaller than this duration would be considered wrong. (default 45s)
  -pg-db-name string
    	PostgreSQL DB name. (default "chicha")
  -pg-host string
    	PostgreSQL DB host. (default "127.0.0.1")
  -pg-pass string
    	PostgreSQL DB password.
  -pg-port int
    	PostgreSQL DB port. (default 5432)
  -pg-ssl string
    	disable / allow / prefer / require / verify-ca / verify-full - PostgreSQL ssl modes: https://www.postgresql.org/docs/current/libpq-ssl.html (default "prefer")
  -pg-user string
    	PostgreSQL DB user. (default "postgres")
  -proxy string
    	Proxy incoming data to another chicha collector. For example: -proxy '10.9.8.7:4000'.
  -race-type string
    	Valid race calculation variants are: 'delayed-start' or 'mass-start'. 1. 'mass-start': start time is not taken into account as everybody starts at the same time, the first gate passage is equal to the short lap, positions are counted based on the minimum time to complete maximum number of laps/stages/gates including the short lap. 2. 'delayed-start': start time is taken into account as everyone starts with some time delay, the first gate passage (short lap) is equal to the start time, positions are counted based on the minimum time to complete maximum number of laps/stages/gates excluding short lap. (default "mass-start")
  -timeout duration
    	Set race timeout duration. After this time if nobody passes the finish line the race will be stopped. Valid time units are: 's' (second), 'm' (minute), 'h' (hour). (default 2m0s)
  -timezone string
    	Set race timezone. Example: Europe/Paris, Africa/Dakar, UTC, https://en.wikipedia.org/wiki/List_of_tz_database_time_zones (default "UTC")
  -version
    	Output version information
  -web string
    	Provide IP address and port to listen for HTTP connections from clients. (default "0.0.0.0:80")
```

## Цветовые подсказки во время гонки:

В авто- и мотоспорте, на соревнованиях, на которых спортсмены борются за лучшее время круга или за наилучший результат в гонке, используется система цветовых сигналов на табло для показа изменений времени круга.

Когда спортсмен завершает круг, его время отображается на табло, и цвет сигнала указывает на то, улучшил ли он свой результат по сравнению с предыдущим кругом или нет. Вот как работает алгоритм:

<span style="color:green">Зеленый цвет:</span> если время круга лучше предыдущего, то на табло будет отображаться зеленый цвет. Это означает, что спортсмен улучшил свой результат, и это может стимулировать его на дальнейшее улучшение времени.

<span style="color:red">Красный цвет:</span> если время круга хуже, чем предыдущее, на табло будет отображаться красный цвет. Это означает, что спортсмен ухудшил свой результат, и ему нужно работать над улучшением.

<span style="color:violet">Фиолетовый цвет:</span> если на табло появляется фиолетовый цвет, это означает, что спортсмен показал лучшее время круга на трассе. Это может быть достигнуто в конце сессии, когда все спортсмены завершают свои круги, или в середине сессии, если спортсмены уже успели улучшить свои результаты.

Цветовые сигналы на табло используются для помощи спортсмену в оценке своей производительности и понимании, насколько он улучшает свои результаты. Это также помогает зрителям понимать, как проходит гонка и кто лидирует.
