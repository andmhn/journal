Android studio
=================

I Installed Android SDK and Android NDK
and in my .bashrc

``` sh

export ANDROID_HOME="~/Android/Sdk/"
export ANDROID_NDK_HOME="~/Android/Sdk/ndk-bundle/29.0.14206865"

```


bundle tool
=================

downloaded from https://github.com/google/bundletool/releases

created a execuable script and put it in my binary path

file: bundletool

``` sh

#! /bin/bash

java -jar ~/Downloads/bundletool-all-1.18.2.jar "$@"

```


Generated key
================
keytool -genkeypair -alias key0 -keyalg RSA -keysize 2048 -validity 365 -keystore keystore.jks


Build Debug
===============
```sh
fyne package -os android -app-id com.github.andmhn.journal -icon icon.png 
```

Release
===============
for relasing to play store
``` sh
fyne release -os android -app-id com.github.andmhn.journal -app-version 1.0 -app-build 1 -icon icon.png  --keystore anand-new.jks --key-name key0
```

for generating signed apk
``` sh
bundletool build-apks --bundle=journal.aab --output=journal_signed.apks --ks=anand-new.jks --ks-key-alias=key0 --mode=universal
```


