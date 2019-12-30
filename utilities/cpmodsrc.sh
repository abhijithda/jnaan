set -x;
# Above 'set -x' should be present, otherwise csh shell complains with the
# below error:
#   server [46]: cpmodsrc root@system.domain.com file.pm
#   function: Command not found.
set +x;

#------------------------------------------------------------------------------#
# cpmodsrc: Copy modified files to a given host and check the syntax on remote
#   host. If host is not specified, then it will only generate the diff. If
#   files are explicitly specified, then only those files will be copied.
#------------------------------------------------------------------------------#


function usage
(
  exit 0;
)


function set_password_less_ssh
(
#my_key="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDUAY/2m6wLuMEMSLX1I629Ocjd4bQ162XE1pyd5L1yPCVxTAOOwWn3doX54lFnkxNKjiv+KkKISLWKIhtp2wKJwXsrZWl2E4v3RWc83hOUzuz/Jx39xdIjtf/OYRMn5DVOMtRdO1APuvXGm3DqE3DHhwU3hN2SEM2vehvVJNkNswf1GPfuoUtElCHfemHwTwHVacAW2bmdLVGX1yim/EqvK3XwmjSSokA+8joMoyEjmb1MaJxRhZzqsMFaEZYgJlQWYqcCSEbNea/rr/JoWUsnJHO/Y9/gi9Xxtz6l1kq0CkEI8VDiDLblHWSJVIcK5bWVcbnWSQw3bvS3y7mQeGUZ root@system.domain.com";
  
  # Specifying BatchMode doesn't prompt for password.
  ssh -o BatchMode=yes -o ConnectTimeout=10 -o StrictHostKeyChecking=no $host "echo ;"
  if [[ $? != 0 ]]
  then 
    my_key=`cat ~/.ssh/id_rsa.pub`;
    ssh -o ConnectTimeout=10 -o StrictHostKeyChecking=no $host "echo $my_key >> ~/.ssh/authorized_keys;";
  fi
)


host=$1;
shift;
if [[ "$host" == "" ]]
then
  echo "Host not specified!";
fi

fcount=$#;
for i in $@
do
  files[${#files[*]}]=$i;
done

echo "Files count " $#;
echo "Files specified " ${files[@]};

#if [[ "$files" == "" ]]
if [[ ${#files[*]} -eq 0 ]]
then
  # cvs diff -uN > ab.diff;
  # for i in `grep -nw Index ab.diff | cut -d ":" -f 3`;

  p4_files=(`~/bin/p4 diff -sa`);
  for i in ${p4_files[@]}
  do
    files[${#files[*]}]=`basename $i`;
  done
  
#if [[ "$files" == "" ]]
  if [[ ${#files[*]} -eq 0 ]]
  then
    echo "No change detected.";
    exit 0;
  fi
  
  echo 
  echo "Detected change in files: ${files[@]}";
  echo 
fi

domain="domain"
if [[ "$host" != "" ]]
then
  
host="root@$host.$domain.com";
echo "Host: $host";

  set_password_less_ssh;

  echo 
  echo "Copying files to $host";
  echo '------------------------';
  # scp "${p4_files[@]}" $host:$SCRIPTS;
  # echo 
  # echo "Syntax check";
  # echo '------------------------';
# ssh $host "cd $SCRIPTS; for i in ${files[@]}; do perl -cw $i; done";

  SCRIPTS="/opt/NBUAppliance/scripts";
  NBA_XML="/opt/NBUAppliance/xml";

  for i in ${p4_files[@]}; do 
    case $i in
      *'.pl'|*'.pm'|*'.t')
        scp $i $host:$SCRIPTS;
        script=`basename $i`
        ssh $host "cd $SCRIPTS; perl -cw $script; chmod a+x $script;";
        ;;
      *'.src')
        scp $i $host:$NBA_XML;
        ssh  $host "cd $NBA_XML; make;";
        ;;
    esac
  done
fi

echo


exit 0;

#------------------------------------------------------------------------------#
fcount=$#;
for i in $@
do
  files[${#files[*]}]=$i;
done

echo "Files count " $#;
echo "Files specified " ${files[@]};


PERFORCE_SOURCE="http://mydesktop.domain.com/code/main/scripts/";
DOWNLOAD_DIR=downloads

mkdir -p $DOWNLOAD_DIR
cd $DOWNLOAD_DIR


for i in ${files[@]};
do
  rm $i;
  wget $PERFORCE_SOURCE/$i;
done
