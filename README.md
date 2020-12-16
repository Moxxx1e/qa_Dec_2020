# qa_Dec_2020

[![Build Status](https://travis-ci.com/Moxxx1e/qa_Dec_2020.svg?branch=develop)](https://travis-ci.com/github/Moxxx1e/qa_Dec_2020)

### Задача №1 
На вход поступает 2 файла:
1) *TestcaseStructure.json*
2) *Values.json*

#### TestcaseStructure.json
Cодержит в себе массив параметров. Каждый **параметр** может иметь поля:
- id - Уникальный идентификатор;
- title - Название параметра;
- value - Выбранное значение, изначально пустая строка;
- values - Массив возможных значений, если поля нет, считаем что значение заполняется вручную, а не выбором.

**Значения** имеют следующие поля:
- id - Уникальный идентификатор;
- title - Название значения, которое подставляется в value параметра;
- params - массив вложенных в это значение параметров.

#### Values.json
Содержит массив объектов. Каждый **объект** имеет поля:
- id - Идентификатор параметра;
- value - Идентификатор значения этого параметра. либо строка, если у параметра нет массива значений.

#### Доп. информация
Если мы не можем найти параметр с id из файла *Values.json*, то считаем, что такого параметра нет и значение подставлять некуда.

Если у параметра с id из *Values.json*, в массиве values нет объекта с id равным value, то считаем, что такого значения нет и оставляем поле value пустым.

Если входные файлы являются неконсистентными, то программа должна сформировать файл *error.json* с сообщением о том, что входные файлы являются некорректными. Пример *error.json* приложен к заданию.

#### Результат
В результате выполнения, программа должна сформировать файл *StructureWithValues.json* с заполненными value у параметров, на основе файла *Values.json*.

Пример корректного выполнения задачи *StructureWithValues.json* приложен к заданию.

Задание принимается только с автотестами(!), уровень и детализацию, которых вы определяете самостоятельно.

### Задача №2
Автомат принимает накопительные скидочные карты и при своем расчете учитывает количество баллов, по которому начисляет процент скидки:
От 0 до 100 баллов - скидка 1%
От 100 до 200 баллов - скидка 3 %
От 200 до 500 баллов - скидка 5%
От 500 баллов -  скидка 10%
Задание: Составить такой набор тестовых данных для автомата, при котором мы гарантированно будем знать, что в соответствии со своими накопленными баллами покупатель получит верную скидку.

#### Результат
Выложить отдельным файлом с названием TaskTestData.md