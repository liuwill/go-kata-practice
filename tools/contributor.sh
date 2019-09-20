#!/bin/sh

printError() {
  echo "\033[31mERROR: $1\033[0m"
}

if [ "$1" == "lines" ];then
  # 某个人的贡献行数
  if [ -n "$2" ]; then
    git log --author="$2" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -
  else
    printError "Missing Username!"
  fi
elif [ "$1" == "count" ];then
  if [ "$2" == "commit" ]; then
    # 总的提交数
    git log --oneline | wc -l
  else
    printError "Not Found!"
  fi
elif [ "$1" == "commit" ]; then
  # 所有人的提交数
  git log --pretty='%aN' | sort | uniq -c | sort -k1 -n -r
else
  # 每个人的贡献行数
  git log --format='%aN' | sort -u | while read name; do echo -en "$name\t"; git log --author="$name" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -; done
fi
