# B.U.P Utility

A utility made with Go that allows you to backup your files with ease.

## Commands

### **bup start**

Intializes **B.U.P** in a folder inputed by the user and creates a config file
containing some useful information about the project.

Example:

```txt
$ bup start
> Backup directory: home/user/Documents/backup_folder
> Project name: example_name
> Project version: example_version
> Author: example_author
> License: example_license
```

After the setup the application should create a config file with the following content:

```json
{
    "ID": "ab530a13e45914982b79f9b7e3fba994cfd1f3fb22f71cea1afbf02b460c6d1d",
    "IsInit": true,
    "CreatedAt": "01-02-2021 17:20",    
    "BackupDirectory": "home/user/Documents/backup_folder",
    "RootDirectory": "home/user/Documents/root_folder",
    "ProjectName": "example_name",
    "ProjectVersion": "0.0.1",
    "Author": "John Doe",
    "License": "M.I.T",
}
```

**Note:** if the folder already exists the program will ask the user to delete it, or end execution

### **bup add**

Saves all the filepaths inside the `root_directory` and stores them in a temporal file to be later used.

Example:

```txt
$ bup add
> Added: /home/user/Documents/root_folder/document.txt
> Added: /home/user/Documents/root_folder/sub_folder/
> Added: /home/user/Documents/root_folder/sub_folder/document2.txt
```

After saving all the filepaths the application should create a temporal file with the following content:

```json
{
    "Filepaths": [
        "/home/user/Documents/root_folder/document.txt",
        "/home/user/Documents/root_folder/sub_folder/",
        "/home/user/Documents/root_folder/sub_folder/document2.txt",
    ],
}

```

**Note:** this command will only work if the config file is already created and initialized.

### **bup commit**

Creates a log file inside both the `root_directory` and the `backup_directory` containing a message inputed by the user.

Example:

```txt
$ bup commit
> Commit message: Example message
> Commit ID:6a31f6b3cd64604c3098 successfully saved within log.json
```

After getting the message the application should create a log.json file with the following content:

```json
[
    {
        "ID": "6a31f6b3cd64604c3098",
        "Message": "Example message",
        "CreatedAt": "01-02-2021 18:15",
    }
]
```

**Note:** this command will only work if the temporal file is already created.

### **bup push**

Copies the files from the `root_directory` to the `backup_directory`

Example:

```txt
$ bup push
> Copying: /home/user/Documents/root_folder/document.txt
> Copying: /home/user/Documents/root_folder/sub_folder/
> Copying: /home/user/Documents/root_folder/sub_folder/document2.txt
> Successfully copied /home/user/Documents/root_folder/ content to /home/user/Documents/backup_folder/
```

After running this command the temporal file will get removed from the `root_directory`

**Note:** this command will only work if the log file is already created.

## TODO's

- [x] Improve the file-structure of the project
- [x] Use packages for code reusability
- [x] Refactor the `start.go` command
- [x] Refactor the `add.go` command
- [x] Refactor the `commit.go` command
- [x] Implement the `push.go` command
- [] Write test(s) for each command
