setlocal
pushd %~dp0

javac *.java
java main

set /p=Hit Enter to Continue
del *.class

popd