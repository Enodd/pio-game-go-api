# PIO-game golang api

Jest to realizacja programu pisanego na zajęciach z podstaw inżynierii oprogramowania. Pierwotnie był on pisany w Java'ie, jednakże postanowiłem zrealizować go w formie api w go.

### Założenia:

- Gracze dołączają przez websocket
- Tworzone jest lobby dla X graczy i Y botów, do którego można dołączyć
- Każdy z graczy w aplikacji internetowej może odgadnąć wynik, boty wykonują rzut automatycznie, gdy przychodzi ich kolej
- Każde lobby będzie miało swoje statystyki
- Historia gier dla danego gracza jest zapisywana w bazie danych.
- Gracze mogą grać jako "Gość" albo utworzyć swoje konta

#### Detale techniczne:

- implementacja dockera w celu łatwiejszego postawienia aplikacji, oraz lepszego DX

