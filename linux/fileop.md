awk 相关应用
========

## 统计日志中出现次数最多的 ip
```bash
awk -F " " '{sum[$1] += 1} END {for(k in sum) print k ":" sum[k]}' filename | sort -n -r -k 2 -t ':' | head -10
-F : 分隔符
$1 : 分隔符分隔后第几个位置（从 1 开始）
```

## 显示文件内容 tail 和 head

```
tail -n 1000 filename：显示最后1000行
tail -n +1000 filename：从1000行开始显示，显示1000行以后的
head -n 1000 filename：显示前面1000行
```

## 显示文件内容 sed 命令
```
sed -n '5,10p' filename 这样你就可以只查看文件的第5行到第10行。
```