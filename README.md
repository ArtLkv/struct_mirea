## Сделано
### Задача 1 - `Selection Sort`
Мы берем первый элемент массива, и бегаем по всему оставшемуся массиву в поисках числа меньше нашего первого элемента. 
Когда мы пробегаем весь массив, если меньшее число найдено — меняем его местами с первым элементом.
### Задача 2 - `Quick Sort`
Cуть алгоритма заключается в разделении массива на два под-массива, средней линией считается элемент, который находится в самом центре массива. 
В ходе работы алгоритма элементы, меньшие чем средний будут перемещены в лево, а большие в право.
### Задача 3 - `Merge Sort` (узнать простая или естественная)
Сначала делим список на кусочки (по 1 элементу), затем сравниваем каждый элемент с соседним, сортируем и объединяем. 
В итоге, все элементы отсортированы и объединены вместе.
### Задача 4 - `Barrier Search`
Сначала меняем искомый элемент на последний в массиве и идем по массиву пока не встретим этот самый элемент(баррьер).
Так можно избавиться от условия сравнения выхода за пределами массива
### Задача 5 - `Binary Search`
Выбираем опорный элемент. Если искомый элемент больше, чем опорный, то элемент находится с право и левую часть можно отбрасывать, или наоборот.
Повторяем пока не найдем искомый. (только с отсортированными массивами)
### Задача 6 - `Simple String Searching`
Смысл задачи в том, чтобы найти подстроку по шаблону в строке. Мы пробегается по словам и ишем, чтобы первый символ шаблона совпадал с символом строки и только за тем идет проверка остальных символов.
### Задача 7.1
Реализовать однонаправленный список со след функциями:
- [X] Узел списка с информационной частью `int`
- [X] Добавление узла(Push)
- [X] Вывод списка(PrintSequence)
Реализация варианта 1:
- [X] Создать на основе алгоритма два списка `L1` и `L2`(наполняем содержимым, я наполнил рандомным)
- [X] Создать на основе алгоритма список `L`, включим в него по одному разу элементы двух предыдущих списков.
### Задача 7.2
Реализовать двунаправленный список со след функциями:
- [X] Узел списка с динамачиеской информационной частью(Для узла списка надо использовать struct)
- [X] Вывод списка в двух направлениях
- [X] Поиск вхождений в списке(возвращает указатель)
- [X] Добавление элемента в список
- [X] Удаление элемента из списка
- [X] Программа должна отдельно выводить выполнения всех требуемых операций(Вывод, Поиск) и выполнять вариант задания
Реализация варианта 1:
- [X] Добавление элемента по номеру зачетной книжки(id)
	- [X] Если такой id имеется, то вставить перед первый элементов с таким же id
	- [X] Если такого id нет, то вставить перед элементом с большим id
	- [X] Если такого id нет и id всех элементов меньше, то ничего не делать 
- [X] Удаление элемента по группе(или по оценке)
- [X] Сформировать новый список из оценок 'неуд' 
## В процессе
### Задача 8
Реализовать стек со следующими функциями:
- [ ] Узел стека с динамачиеской информационной частью
- [ ] Добавить элемент в стек
- [ ] Убрать элемент из стека
- [ ] Получить значение узла стека
- [ ] Проверить полон ли стек
- [ ] Проверить пуст ли стек
Реализация варианта 2:
