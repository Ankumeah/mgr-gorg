# mgr-gorg
Simple program in golang to download all your favorite mangas on [mangaread](https://mangaread.org)

## How to build
```sh
go build -ldflags="-s -w" -o ./bin/mgr-gorg ./src/
```

## How to use
- Go to [mangaread](https://mangaread.org)
- Navigate to the manga you want to download (*for example - https://www.mangaread.org/manga/we-never-learn/*)
- Copy the last part of the URL (*in this example we-never-learn*)
- Run the command `mgr-gorg <your-manga> <chapters you want>` (in this example `mgr-gorg we-never-learn 1 2 3 4 5`)
- A folder with the same name as your manga will appear in your current working directory with the individual jpgs themselves along with their respective .cbz files

> Note: This project is very much untested so feel free to open any issues if you find any or if you are willing then you can even open a pull request, all (except AI) are welcome :)
