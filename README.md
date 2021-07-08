# helper

golang cli helper: personal utilities

## Features (current)

* *Quick* HTTP static fileserver

----------

## Features (todo)

* **ssh command**
  * configure subcommand that saves info into json file (load pubkeys, set hosts)

* **linux user management**
  <br>
   \- *all capitalized words prefixed with 1 dollarsign are mandatory*
  <br>
  \- *all capitalized words prefixed with 2 dollarsigns are optional*
  * short alias for:
    * sudo `usermod` -aG $GROUP $USER

    * sudo `useradd` --badnames -omU -u 1000 -p $$PASSWORD -s $$SHELL $USERNAME
      * **execute this cmd after**:
        * `mkdir` -p ~/.cache/oh-my-zsh

* **system initialization menu**
  <br>
  initialize things such as
  * **github credentials** (github-cli && git [user. name && user.email])
  * insert **~/.zshrc** for both the user *root* and the *active* user
