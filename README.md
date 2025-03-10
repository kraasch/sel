
# sel

A small program to make selections:

```bash
echo -e "melon juice\ntoast\ngrapes" | sel -verbose
```

## tasks

  - feature: implement undo-redo history.
    - save a list of the last indices which changed.
    - make the U key undo that list.
    - make `ctrl-r` redo these changes.
    - after going back with u key and making changes with other keys, chop off history which was undone and make new keys start a new history from current point.
  - bug: this freezes the terminal completely, terminal has to be killed:
    - make sure usage in shell sub-commands is non-blocking.
    - blocking example: `echo $(echo -e "melon juice\ntoast\ngrapes" | sel)`
  - bug: fix bug #1

## bugs #1

  - boxes not aligned (front misaligned).
  - endings print double ("toastst") and fill the space lost in the front.

```bash
~/repos/remote_visible/dev/sel$ echo -e "melon juice\ntoast\ngrapes" | sel | while read x; do echo "$x"; done
> [ ] melon juice
[ ] toastst
[ ] grapeses
```


