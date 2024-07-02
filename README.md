# GoTutorial

This repo contains code that I wrote while learning Go. I decided to learn Go mainly to see what all the hype around it is. I know that it's supposedly very performant but unfortunately many people say that its not the easiest code to write and might be more C like. Will I like Go by the end of this journey? Maybe not but it will definetly easier for me to know Go when the time comes for me to need to write it cause the signs of a project refactor using Go are there sadly.
  
<img src="https://media1.tenor.com/m/XXyXCrGg6nkAAAAC/hyunrmin.gif" alt="crying" width=500 height=600>

Very amazing beginning to the project I've already messed up the modules! It seems like this directory should have been a module hopefully its an easy fix!

```bash
go mod init <module-name>
```

So a module is a collection of packages and a package is a folder containing GO files. Sounds easy enough. Okay maybe I didn't fuck up! I thought creating a new module would create a new folder in a similar way creating a new Node project creates a new folder but it looks like Go doesn't do that and instead just creates a **go.mod** file. Looks like it specifies the module name and go version. Seems to be similar to something like a package.json. Similar to package.json it also has details of installed modules.

All your code will be contained inside packages. These packages are contained inside folders. All the functionality / files that belong to the same package should be in the same folder. 

Go code requires an entrypoint. The entrypoint is defined in the main package. The main package is any package whose name is main. BTW packages should have names. IDK if this is a must but maybe a package should have a main.go file. This file includes the package name at top defined by doing `package <package-name>`. The main package must have a main function which I think must be defined in the main.go file?. 
