# core

First, make sure those components are install properly on your machine:  
- Golang (https://golang.org/doc/install)
- MySql Server (you can use xampp with will provide your mysqlserver and a phpmyadmin https://www.apachefriends.org/fr/download.html)
- Git (https://git-scm.com/download/win)
- I also advice you to use visual studio CODE which is a light and powerful text editor :)

If everything works fine, you should be able to run the application *git bash* and run the command go.  

## Get the DB ready

If you installed xampp, you can open phpmyadmin (or any other equivalent) and create a new database called `yana`.
When the database has been created, you can run the script `core/_resources/create.sql`. It will create all required tables.

Then you can run `core/_resources/fill_db.sql` in order to insert mandatory rows.

Last step before having the database ready. You will need to update the configuration file of Yana in order to use your own credentials:  
On yana-app & api you can find the file env.json. On this file, please replace the database settings:
```
"MySQL": {
		"Username": "{YOUR SQL USERNAME}",
		"Password": "{YOUR SQL DEFAULT PASSWORD}",
		"Database": "yana",
		"Charset": "utf8mb4",
		"Collation": "utf8mb4_unicode_ci",
		"Hostname": "127.0.0.1",
		"Port": 3306,
		"Parameter": "parseTime=true",
		"Migration": {
			"Folder": "migration/mysql",
			"Table": "migration_yana",
			"Extension": "sql"
		}
	},
```

## Download and run Yana

First, create a folder called *yana* in your GOFOLDER/src/  
Then, simply download core, api, etc.. and put them in this exact folder.  

Your folders should be link this:
  GOFOLDER/src/yana  
  					/yana-app  
					/api  
					/core  
					
Go should install automatically the dependancies. If not: https://coderwall.com/p/arxtja/install-all-go-project-dependencies-in-one-command

To run the project, you have 2 ways:
- using git bash, you can do:
```
$ cd {yourgofolder}/src/yana/api
$ go run main.go
```
- using file explorer, navigate to your go folder and then to: src, github.com, yana-bar, api. Then right click and use *Open git bash here*. Then run:
```
$ go run main.go
```

The application should listen to localhost:8080 so if you look for `http://localhost:8080` on your navigator, you should see a documentation page.

## Project organization

The aim of using go is to split the logic and the application properly. core contains the logic to get the data (get articles, get product, create new order, update order, etc...). api and yana-app (and soon yana-landing) will simply call the core methods to do their job.

## Git

Git is a powerful tool which allows us to work together on the project. In order to avoid any issue, please create your own branch.  
To do so, go on core folder and type
```
$ git checkout -b {typeTheNameYouWant}
```
If everything's fine, it should display a new branch. You will have to repeat the same action for the 3 folder (core, api, yana-app)

