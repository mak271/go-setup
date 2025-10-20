package main

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	//"strconv"
	//"sync"
)

type Product struct {
	Name  string
	Price int
}

type Order struct {
	CustomerID int
	Price      int
}

func SortOrdersByCustomerID(orders []Order) []Order {
	ordersCopy := slices.Clone(orders)
	slices.SortFunc(ordersCopy, func(first, second Order) int {
		if first.CustomerID == second.CustomerID {
			return first.Price - second.Price
		}
		return first.CustomerID - second.CustomerID
	})
	return ordersCopy
}

func NewDiscountedProduct(name string, price int, discount int) *Product {
	mPrice := price
	if discount > 0 {
		mPrice = price - (price * discount / 100)
	}
	if discount > 100 {
		mPrice = 0
	}
	product := Product{Name: name, Price: mPrice}
	return &product
}

func AreOrderHistoriesEqual(history1, history2 [][]string) bool {
	// Явно проверяем: один nil, другой пустой
	if (history1 == nil) != (history2 == nil) {
		return false
	}
	if len(history1) != len(history2) {
		return false
	}
	for i := range history1 {
		if !slices.Equal(history1[i], history2[i]) {
			return false
		}
	}
	return true
}

func CompareProductLists(oldList, newList []string) (added, removed []string) {
	for _, v := range newList {
		if !slices.Contains(oldList, v) {
			added = append(added, v)
		}
	}
	for _, v := range oldList {
		if !slices.Contains(newList, v) {
			removed = append(removed, v)
		}
	}
	return added, removed
}

func GetGrade(scores map[string]int, name string) (string, error) {
	if scores[name] == 0 {
		return "", errors.New("Error")
	}
	return fmt.Sprintln(name, "has", scores[name], "points"), nil
}

func CountWords(text string) map[string]int {
	result := make(map[string]int)
	words := strings.FieldsSeq(strings.ToLower(text))

	for word := range words {
		// Удаляем знаки пунктуации с конца и начала
		word = strings.Trim(word, ".,!?;:")
		if word != "" {
			result[word]++
		}
	}
	return result
}

type User struct {
	Name  string
	Email string
}

func UpdateEmail(users map[int]*User, id int, newEmail string) error {
	user, ok := users[id]
	if !ok {
		return errors.New("user not found")
	}
	user.Email = newEmail
	return nil
}

func SetUserSetting(settings map[string]map[string]string, user, key, value string) {
	if settings[user] == nil {
		settings[user] = make(map[string]string)
	}
	settings[user][key] = value
}

func CountLanguages(users map[string]string) map[string]int {
	languages := map[string]int{}
	for _, language := range users {
		languages[language]++
	}
	return languages
}

func main() {
	// BEGIN (write your solution here)
	// wg := sync.WaitGroup{}

	// for i := range 3 {
	// 	wg.Go(func() {
	// 		fmt.Println("Go! " + strconv.Itoa(i))
	// 	})
	// }
	// wg.Wait()
	// END
	users := map[int]*User{
		1: {Name: "Alice", Email: "alice@example.com"},
		2: {Name: "Bob", Email: "bob@example.com"},
	}

	err := UpdateEmail(users, 1, "alice@newmail.com")
	fmt.Println(users[1].Email) // "alice@newmail.com"
	fmt.Println(err)            // <nil>

	err = UpdateEmail(users, 3, "charlie@mail.com")
	fmt.Println(err) // "user not found"
}

//fmt.Sprintf("Package %s has been delivered", p.ID)
//fmt.Sprintf("%.1f", rating)
//fmt.Printf("%s is %d years old\n", name, age)
//fmt.Println("Port:", cfg.Port)
// var n int = 65
// fmt.Println(strconv.Itoa(n)) // => "65"
// var s string = "42"
// n, err := strconv.Atoi(s)
// fmt.Println(n) // => 42
// Создаём копию только нужной части среза с помощью slices.Clone
//return slices.Clone(src[:maxLen])

// buf := []int{1,2,3,4,5}
// newBuf = make([]int, len(buf), len(buf))
// copy(newBuf, buf)

// 	fmt.Printf("%v\n", files)
