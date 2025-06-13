# GTH: Go Templ Htmx 

A Starter kit made in a evening for people hate writing javascript and I call this GTH because every tech stack needs a bad name.

The tech stack is 
- Go: for the back end 
- Gin: A library for handling route 
- Templ: To handle templates and html 
- Tailwind: Because CSS is the worst, and htmx and tailwind works hand in hand 
- Htmx: To handle frontend code with 0 java script
- Optionally Alpine.js: for more interactivity 

## Getting Started 
To install this starter kit:

`git clone https://github.com/ThatAdwaithGuy/gth-starterkit.git && cd gth-starterkit`

after that run 

`./install.sh *YOUR GITHUB USERNAME*/*YOUR PROJECT NAME*`

Your project name is used as the go's module name, That's why your github username is needed for publishing. I am kinda new to go so please help me in this matter.

Then rename "gth-starterkit" to your projects name.

If your project doesn't need a database, 
- Run `./deletedb.sh`
- Use `make generate`

If your app needs a database, 
- Make a .env file and create DATABASE_URL variable and connect it with your database(REQUIRED).
- By Default, the starterkit uses postgres (Best database in my opinion) but you can change it. Please visit sqlc's website and check out the other options.
- Uncomment code in the main.go file (Instructions are in the file)
-  Use `make build`


Then run the suitable make generate command.

Use `air` to run the app with live reloading

## Requirement
- [air](https://github.com/air-verse/air)
- [sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)
- [dbmate](https://github.com/amacneil/dbmate)

From here, everything is free to the imagination. Have fun using this :).




