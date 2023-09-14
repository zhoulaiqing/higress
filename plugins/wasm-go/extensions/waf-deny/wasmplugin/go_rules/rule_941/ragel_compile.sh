for file in `ls *.rl` ; do
  echo "Compiling: "  $file
  ragel -Z $file
done