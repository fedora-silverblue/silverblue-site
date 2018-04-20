# Team Silverblue
[Team Silverblue](https://teamsilverblue.org) main website. This is a simple static HTML generator that can be run locally, provided that Go is installed. A CI environment will be added at a later stage.

## Quickstart
front-end: HTML5, CSS

back-end: Go

### How To Run
From your terminal, navigate to the repository and type the following commands for a local server.

```
cd src
go run main.go
```

Go to [localhost:8080](http://localhost:8080) in your browser and you will see the website.

For production use, for example as your blog engine, it's recommended to use `go install` and then run the built executable that is located in the `bin`
folder of the repository.

### How To Contribute
Team Silverblue wants to make it possible for you to easily contribute to the docs by sharing finished content as well as ideas - if you want to share them here, please create a pull request and use *whichever text format works best for you* by dropping the file into the `ideas` folder using a descriptive file name. We might ask you to elaborate on your idea/suggestion in an IRC meeting. If all that is missing in your file are HTML tags or it needs a conversion of any text file format into HTML, we will convert it into HTML and move it to `/pages` when it is edited and approved. This process should enable technical writers as well as developers and ops engineers to contribute to docs for Team Silverblue without starting hurdles.

For finished content, please create a file with the following naming convention: `yyyy-mm-dd-title-slug.html`. The provided date and title slug will also be the URL for your blog title without the appended `.html`. You'll have to be somewhat proficient in HTML. This website uses meaningful HTML5.

The structure of the code is as follows. Inline comments in the files contain details about each file.

```
/ideas          # ideas, suggestions, and raw files in other text formats go here
/pages          # content in between header and footer goes here
/public         # images and all files that need to be statically served go here
/src            # the actual server and templating engine
/templates      # main structure and layout files
```

If you have questions or want to contribute but don't know where to start, ask us via Twitter [@teamsilverblue](https://twitter.com/teamsilverblue).
