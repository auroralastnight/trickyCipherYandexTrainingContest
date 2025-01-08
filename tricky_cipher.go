package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
  "unicode/utf8"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  n, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println("Ошибка ввода:", err)
  }

  iterationN, _ := strconv.Atoi(strings.TrimSpace(n))

  var userInput string
  var value []string
  userInfo := map[string]string{
    "surname":    " ",
    "name":       " ",
    "patronymic": " ",
    "dayBorn":    " ",
    "monthBorn":  " ",
    "yearBorn":   " ",
  }
  for range iterationN {
    userInput, _ = reader.ReadString('\n')
    userInput = strings.TrimSpace(userInput)
    value = strings.Split(userInput, ",")
    userInfo["surname"] = value[0]
    userInfo["name"] = value[1]
    userInfo["patronymic"] = value[2]
    userInfo["dayBorn"] = value[3]
    userInfo["monthBorn"] = value[4]
    userInfo["yearBorn"] = value[5]

    namesCharCount := countUniqueLetters(userInfo["surname"], userInfo["name"], userInfo["patronymic"])
    sumBDayNum := sumBirthdayIntegers(userInfo["dayBorn"], userInfo["monthBorn"])
    firstLatterNum := firstLetter(userInfo["surname"])
    answer := hexCode(namesCharCount, sumBDayNum, firstLatterNum)

    fmt.Print(answer, " ")
  }
}

func countUniqueLetters(surname string, name string, patronymic string) int {
  charCounter := make(map[rune]bool)
  count := 0
  for _, char := range name + surname + patronymic {
    if charCounter[char] {
      continue
    }
    charCounter[char] = true
    count++
  }
  return count
}

func sumBirthdayIntegers(day string, month string) int {
  sum := 0
  var num int
  for _, char := range day + month {
    num, _ = strconv.Atoi(string(char))
    sum += num
  }
  return sum
}

func firstLetter(surname string) int {
  return int(surname[0]) - 64
}

func hexCode(namesCharCount int, sumBDayNum int, firstLetterNum int) string {
  hexadecimal := fmt.Sprintf("%X", (namesCharCount + sumBDayNum*64 + firstLetterNum*256))
  length := utf8.RuneCountInString(hexadecimal)
  if length >= 3 {
    return hexadecimal[length-3:]
  }
  for length < 3 {
    hexadecimal = "0" + hexadecimal
    length++
  }
  return hexadecimal
}
