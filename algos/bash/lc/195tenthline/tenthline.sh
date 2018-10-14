# Read from the file file.txt and output the tenth line to stdout.
file=${1:-file.txt}
# cat -n file.txt | grep 10 | cut -d "0" -f 2-
echo "Filename: $file"
cat -n ${file} | grep -w 10 | cut -c 8-