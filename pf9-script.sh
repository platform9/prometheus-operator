#!/usr/bin/env bash
eval "$(gimme 1.19)"
export PATH="$PATH:$HOME/go/bin"
make pf9-image || exit
make pf9-push || exit
