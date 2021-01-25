# B.U.P Backup utility made with Go
A simple tool that allows you to backup and keep track of your projects

## B.U.P Commands in detail
### - bup start
Initializes **B.U.P** in a specific directory taking the following parameters from the user:

```bash
$ bup start
> Backup destination: example_route
> Project name: example_name
> Project version (0.0.1) by default: example_version
> Author (John Doe) by default: example_author
> License (M.I.T) by default: example_license
```

**¿How would it work?**
First of all it should store the data inputed by the user, after that it should check
the data, particularly the directory where the backup will occur, if it does exists, the program
will end execution, and if it doesn't exist the program will create that directory. After successfully
creating the dir we should store all the already verifed data inside a newly created config.json file and
also create a log.json that will hold the changes that happen to that specific directory.

### - bup add
Saves the paths of the files within the working folder to a temporal .json

```bash
$ bup add
> Added: example_dir/example_file.txt
> Added: example_dir/example_subdir
> Added: example_dir/example_subdir/example_subfile.txt
> Filepath(s) successfully added to temp_files.json... exit code 0
```

**¿How would it work?**
The program will loop through the working directory to get the paths of the files and the folders, after looping through the dir, those paths will be stored in a temporal .json file

## TODO's
* Write the tests for **add/add.go**
* Allow the user to input specific files in **bup add** command