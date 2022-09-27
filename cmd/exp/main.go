package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgressConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgressConfig) string() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgressConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.string())

	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected:...")
	//Create users,tweets and likes tables
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (  
	id SERIAL PRIMARY KEY,
	name TEXT,
	email TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS tweets ( 
		id SERIAL PRIMARY KEY,
		user_ID INT NOT NULL,
		tweet TEXT
	);

	CREATE TABLE IF NOT EXISTS likes ( 
		id SERIAL PRIMARY KEY,
		user_ID INT NOT NULL,
		tweet_ID INT,
		likes INT 
		);

	`)

	if err != nil {
		panic(err)
	}
	fmt.Println("Tables Created")

	// //INSERT INTO USERS TABLE
	// name := "mohammed waziri"
	// email := "moh@gmail.com"
	// row := db.QueryRow(`
	// 	INSERT INTO users (name,email)
	// 	VALUES ($1,$2) RETURNING id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User created! ID=", id)

	// INSERT INTO TWEET TABLE
	// user_id := 1
	// for i := 1; i <= 3; i++ {
	// 	tweet := fmt.Sprintf("fake tweet %d", i)
	// 	_, err := db.Exec(`
	// 	INSERT INTO tweets (user_ID, tweet)
	// 	VALUES ($1,$2);`, user_id, tweet)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("tweet inserted id=", user_id)

	// user_id = 3
	// for i := 1; i <= 3; i++ {
	// 	tweet := fmt.Sprintf("this is my %d tweet", i)
	// 	_, err := db.Exec(`
	// 	INSERT INTO tweets (user_ID, tweet)
	// 	VALUES ($1,$2);`, user_id, tweet)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("tweet inserted id=", user_id)

	//INSERT INTO likes TABLE
	// user_id := 3
	// tweet_id := 5
	// for i := 1; i <= 8; i++ {
	// 	_, err := db.Exec(`
	// 	INSERT INTO likes (user_ID,tweet_ID,likes)
	// 	VALUES ($1,$2,$3);`, user_id, tweet_id, i)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// fmt.Println("Tweet Liked!")

	// SELECT A USER
	// user_id := 3
	// row := db.QueryRow(`
	// SELECT name,email FROM users
	// WHERE id=$1;`, user_id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("User information name=", name, "email=", email)

	//SELECTING MULTIPLE ROWS ON OUR TWEETS TABLE
	type Tweets struct {
		id      int
		user_id int
		tweet   string
	}

	var tweets []Tweets
	user_id := 3
	rows, err := db.Query(`
	SELECT id,tweet FROM tweets
	WHERE user_id=$1;`, user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var tweet Tweets
		tweet.user_id = user_id
		err := rows.Scan(&tweet.id, &tweet.tweet)
		if err != nil {
			panic(err)
		}
		tweets = append(tweets, tweet)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("tweets:", tweets)

	// //Create a table
	// _, err = db.Exec(`
	// CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT,
	// 	email TEXT UNIQUE NOT NULL
	// );

	// CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	// );
	// `)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Tables Created")

	// // // Insert some data
	// // name := "Hayatu waziri"
	// // email := "bawazy@gmail.com"
	// // row := db.QueryRow(`
	// // 	INSERT INTO users (name,email)
	// // 	VALUES ($1, $2) RETURNING id;`, name, email)
	// // var id int
	// // err = row.Scan(&id)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Println("User created! id= ", id)
	// // id := 1
	// // row := db.QueryRow(`
	// // 	SELECT name,email FROM users WHERE id=$1;`, id)
	// // var name, email string
	// // err = row.Scan(&name, &email)
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Printf("User information: name=%s, email=%s\n", name, email)
	// // userId := 1
	// // for i := 1; i <= 5; i++ {
	// // 	amount := i * 100
	// // 	desc := fmt.Sprintf("Fake Order #%d", i)
	// // _, err := db.Exec(`
	// // INSERT INTO orders(user_id,amount,description)
	// // VALUES($1,$2,$3);`, userId, amount, desc)
	// // 	if err != nil {
	// // 		panic(err)
	// // 	}
	// // }
	// // fmt.Println("Added fake orders")

	// type Order struct {
	// 	Id          int
	// 	User_id     int
	// 	Amount      string
	// 	description string
	// }
	// var orders []Order
	// userId := 1
	// rows, err := db.Query(`
	// SELECT id,amount,description
	// FROM orders
	// WHERE user_id=$1`, userId)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var order Order
	// 	order.User_id = userId
	// 	err := rows.Scan(&order.Id, &order.Amount, &order.description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Orders:", orders)
}
