# 1. Chat

## Napisać aplikację typu chat (5 pkt.)

- Klienci łączą się serwerem przez protokół TCP
- Serwer przyjmuje wiadomości od każdego klienta i rozsyła je do pozostałych (wraz z id/nickiem klienta)
- Serwer jest wielowątkowy - każde połączenie od klienta powinno mieć swój wątek
- Proszę zwrócić uwagę na poprawną obsługę wątków

## Zadanie domowe c.d. (3 pkt.)

- Dodać dodatkowy kanał UDP
- Serwer oraz każdy klient otwierają dodatkowy kanał UDP (ten sam numer portu jak przy TCP)
- Po wpisaniu komendy ‘U’ u klienta przesyłana jest wiadomość przez UDP na serwer, który rozsyła ją do pozostałych klientów
- Wiadomość symuluje dane multimedialne (można np. wysłać ASCII Art)

## Zadanie domowe c.d. (2 ptk.)

- Zaimplementować powyższy punkt w wersji multicast
- Nie zamiast, tylko jako alternatywna opcja do wyboru (komenda ‘M’)
- Multicast przesyła bezpośrednio do wszystkich przez adres grupowy (serwer może, ale nie musi odbierać)

## Zadanie domowe - uwagi

- Zadanie można oddać w dowolnym języku programowania
- Nie wolno korzystać z frameworków do komunikacji sieciowej - tylko gniazda! Nie wolno też korzystać z Akka
- Przy oddawaniu należy:
- zademonstrować działanie aplikacji (serwer + min. 2 klientów)
- omówić kod źródłowy
- Proszę zwrócić uwagę na: wydajność rozwiązania (np. pula wątków), poprawność rozwiązania (np. unikanie wysyłania wiadomości do nadawcy, obsługa wątków)