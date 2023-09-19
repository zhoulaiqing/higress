for file in `ls *.rl` ; do
  echo "Compiling: "  $file
  ragel -Z -G2 $file
done