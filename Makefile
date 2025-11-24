build_dir      =  build
name           =  journal
debug_bin      =  $(build_dir)/journal_debug
release_bin    =  $(build_dir)/journal_release
linux_package  =  $(build_dir)/journal.tar.xz
android_apk    =  $(build_dir)/journal.apk
android_aab    =  $(build_dir)/journal.aab
temp_apks      =  $(build_dir)/temp.apks
signed_apk     =  $(build_dir)/journal_signed.apk

keystore_file  =  ~/code/keystore.kjs
key_alias      =  key0

all:
	go build -o $(debug_bin)

init:
	mkdir -p $(build_dir)

release: init
	go build -o $(release_bin) -ldflags="-w" 

linux-dist: init
	fyne package -os linux -icon icon.png -release
	mv $(name).tar.xz $(linux_package)

android: init
	fyne package -os android -app-id com.github.andmhn.journal -icon icon.png
	mv $(name).apk $(android_apk)

android-release: init
	-rm -f $(temp_apks)
	-rm -f $(android_aab)
	@read -p "Enter KeyStore Password : " STORE_PASS; \
		fyne release -os android -app-id com.example.myapp -app-version 1.0 -app-build 1 -icon icon.png  --keystore $(keystore_file) --key-name $(key_alias) --keystore-pass $$STORE_PASS; \
		mv $(name).aab $(android_aab); \
		bundletool build-apks --bundle=$(android_aab) --output=$(temp_apks) --ks=$(keystore_file) --ks-key-alias=$(key_alias) --ks-pass=pass:$$STORE_PASS  --mode=universal
	unzip $(temp_apks)
	mv universal.apk $(signed_apk)
	rm toc.pb
	-rm -f $(temp_apks)

clean:
	-rm -f -r $(build_dir) toc.pb  universal.apk $(wildcard *.exe) $(wildcard *.apk) $(wildcard *.apks) $(wildcard *.aab) $(wildcard *.pb) 
