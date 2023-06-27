# sem06-distributed-systems

## Oceny

| Laboratorium | Data | Prowadzący | Zdobyte pts | Max pts |
|---|---|---|---|---|
| 1 | 14.03.2023 | Straś | 8 | 10 |
| 2 | 21.03.2023 | Konieczny | 10 | 10 |
| 3 | 25.04.2023 | Konieczny | 6 | 10 |
| 4 | 09.05.2023 | Czekierda | 11.5 | 12 |
| 5 | 16.05.2023 | Czekierda | 4 | 8 |
| 6 | 13.06.2023 | Straś | 10 | 10 |
| 7 | 20.06.2023 | Nawrocki | 7 | 10 |

## Technologie

### Lab01

* TCP / UDP
* websockets
* Go

### Lab02

* Rest API
* Go

### Lab03

* Ray framework
* Python3.9 (Python musi być mniejszy niż 3.10)

### Lab04

* ICE
* Java (server)
* Python (client)

### Lab05

* gRPC
* Go (server)
* Python (client)

### Lab06

* RabbitMQ
* Go

### Lab07

* Zookeeper
* Python

## Opisy

### Lab01

Raczej przyjemne, lekko oceniane

### Lab02

Chwalił za Go, w sumie też spoko, nie polecam Pythona - mnie pokonał

### Lab03

Trzeba pamiętać jak się Ray'a odpala, w zasadzie pytał na ile punktów się robiło i dawał coś pośrodku

### Lab04 i Lab05

Chuck cisnął z pytaniami o wymagania ale jak się mówiło samemu sporo od siebie to raczej na luzie oceniał - jak na to że dostałem 11.5/12 pts mimo brakujących wielu wymagań to nawet nieźle \
Warto się też z raportu przygotować i go napisać a nie tylko ułomne i nie kompletne logi z Wiresharka bo za to dał (aż) 4/8 pts

### Lab06

Powiedział, że spoko, że w Go. Podobało mu się rozbicie kodu na moduły i konfiguracje - od strzała dał 10/10 nie zadając zbyt wiele pytań.

### Lab07

Dr Nawrocki "nie miał czasu" na ocenianie lub był bardzo zmęczony maratonem ocen. Na szybko kazał włączyć 3 serwery, 3 klientów i podawał komendy co należy wpisać w którym kliencie i patrzył czy apka reaguje. Był wielce niezadowolny, że "drzewo" węzłów to po prostu coś co przypomina wyżyg polecenie `ls -Ral` a nie jest graficznie ładnym drzewem. Słusznie obciął punkty za złe zliczanie potomków (nie zliczało rekursywnie - wnuki, prawnuki … - a tylko dzieci zadanego węzła)
