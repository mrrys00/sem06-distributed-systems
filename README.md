# sem06-distributed-systems

## 1. Zadanie domowe - Chat

### Napisać aplikację typu chat (5 pkt.)

- Klienci łączą się serwerem przez protokół TCP
- Serwer przyjmuje wiadomości od każdego klienta i rozsyła je do pozostałych (wraz z id/nickiem klienta)
- Serwer jest wielowątkowy - każde połączenie od klienta powinno mieć swój wątek
- Proszę zwrócić uwagę na poprawną obsługę wątków

### Zadanie domowe c.d. (3 pkt.)

- Dodać dodatkowy kanał UDP
- Serwer oraz każdy klient otwierają dodatkowy kanał UDP (ten sam numer portu jak przy TCP)
- Po wpisaniu komendy ‘U’ u klienta przesyłana jest wiadomość przez UDP na serwer, który rozsyła ją do pozostałych klientów
- Wiadomość symuluje dane multimedialne (można np. wysłać ASCII Art)

### Zadanie domowe c.d. (2 ptk.)

- Zaimplementować powyższy punkt w wersji multicast
- Nie zamiast, tylko jako alternatywna opcja do wyboru (komenda ‘M’)
- Multicast przesyła bezpośrednio do wszystkich przez adres grupowy (serwer może, ale nie musi odbierać)

### Zadanie domowe - uwagi

- Zadanie można oddać w dowolnym języku programowania
- Nie wolno korzystać z frameworków do komunikacji sieciowej - tylko gniazda! Nie wolno też korzystać z Akka
- Przy oddawaniu należy:
- zademonstrować działanie aplikacji (serwer + min. 2 klientów)
- omówić kod źródłowy
- Proszę zwrócić uwagę na: wydajność rozwiązania (np. pula wątków), poprawność rozwiązania (np. unikanie wysyłania wiadomości do nadawcy, obsługa wątków)


## RESTful API

Celem zadania jest napisanie prostego serwisu webowego realizującego pewną złożoną funkcjonalność w oparciu o otwarte serwisy udostępniające REST API. Stworzyć macie Państwo serwis, który:

- udostępni klientowi statyczną stronę HTML z formularzem do zebrania parametrów żądania,
- odbierze zapytanie od klienta,
- odpyta serwis publiczny (różne endpointy), a lepiej kilka serwisów o dane potrzebne do skonstruowania odpowiedzi,
- dokona odróbki otrzymanych odpowiedzi (np.: wyciągnięcie średniej, znalezienie ekstremów, porównanie wartości z różnych serwisów itp.),
- wygeneruje i wyśle odpowiedź do klienta (statyczna strona HTML z wynikami).

Wybranie realizowanej funkcjonalności i używanych serwisów pozostawiam Państwa wyobraźni, zainteresowaniom i rozsądkowi. Przykładowo:
Klient podaje miasto i okres czasu ('daty od/do' lub 'ostatnie n dni'), serwer odpytuje ogólnodostępny serwis pogodowy o temperatury w poszczególne dni, oblicza średnią i znajduje ekstrema i wysyła do klienta wszystkie wartości (tzn. prostą stronę z tymi danymi). Ewentualnie serwer odpytuje kilka serwisów pogodowych i podaje różnice w podawanych prognozach.
Z reguły wygrywa prognoza pogody lub kursy walut, ale liczę na wykazanie się większą kreatywnością ;-) Listę różnych publicznych API można znaleźć np.: na https://publicapis.dev/

### Wymagania

- Klient (przeglądarka) ma wysyłać żądanie w oparciu o dane z formularza (statyczny HTML) i otrzymać odpowiedź w formie prostej strony www, wygenerowanej przez tworzony serwis. Wystarczy użyć czystego HTML, bez stylizacji, bez dodatkowych bibliotek frontendowych (nie jest to elementem oceny). Nie musi być piękne, ma działać.
- Tworzony serwis powinien wykonać kilka zapytań (np.: o różne dane, do kilku serwisów itp.). Niech Państwa rozwiązanie nie będzie tylko takim proxy do jednego istniejącego serwisu i niech zapewnia dodatkową logikę (to będzie elementem oceny).
- Odpowiedź dla klienta musi być generowana przez serwer na podstawie: 1) żądań REST do publicznych serwisów i 2) lokalnej obróbki uzyskanych odpowiedzi.
- Serwer ma być uruchomiony na własnym serwerze aplikacyjnym działającym poza IDE (lub analogicznej technologii).
- Dodatkowym (ale nieobowiązkowym) atutem jest wystawienie serwisu w chmurze (np.: Heroku). To jest część dla zainteresowanych i nie podlega podstawowej ocenie.
- Dopuszczalna jest realizacja zadania w dowolnym wybranym języku/technologii (oczywiście sugerowany jest Python i FastAPI). Proszę jednak o zachowanie analogicznego poziomu abstrakcji (operowanie bezpośrednio na żądaniach/odpowiedziach HTTP, kontrola generowania/odbierania danych).
- Implementacja autoryzacji jest elementem oceny.
- Wybieramy serwisy otwarte lub dające dostęp ograniczony, lecz darmowy, np.: używające kodów deweloperskich.
- Dodatkowo (jest to elementem oceny): Przygotowujemy test zapytań HTTP z wykorzystaniem POSTMANa/SwaggerUI (klient-serwer i serwer-serwis_publiczny). Do oddania proszę mieć je już zapisane.

### Na co warto zwrócić uwagę?

- (!!) obsługę (a)synchroniczności zapytań serwera do serwisów zewnętrznych (np.: promises),
- (!) obsługę błędów i odpowiedzi z serwisów (np.: jeśli pojawi się błąd komunikacji z serwisem zewnętrznym, to generujemy odpowiedni komunikat do klienta, a nie 501 Internal server error),
- walidację danych wprowadzanych przez klienta/przyjmowanych przez serwer.

### Punktacja

- Implementacja serwera - obsługa zapytań do zewnętrznego serwisu: [0-2] pkt.
- Implementacja serwera - odbiór żądań klienta, generowanie i wysłanie odpowiedzi: [0-2] pkt.
- Implementacja serwera - obsługa asynchroniczności zapytań i błędów [0-3]
- Implementacja klienta - statyczny formularz zapytań / strona odpowiedzi: [0-2] pkt.
- Testowanie żądań REST z pomocą Postman-a/Swager UI (do serwera i do serwisu zewnętrznego): [0-1] pkt
