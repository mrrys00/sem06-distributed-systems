.PHONY: lab01_zip
lab01_zip:
	mkdir Rys_Szymon_1
	mkdir -p archives
	cp -r lab01/ Rys_Szymon_1/
	zip -r archives/Rys_Szymon_1.zip Rys_Szymon_1/ oswiadczenie.txt
	rm -r Rys_Szymon_1

.PHONY: lab02_zip
lab02_zip:
	mkdir Rys_Szymon_2
	mkdir -p archives
	cp -r lab02/ Rys_Szymon_2/
	zip -r archives/Rys_Szymon_2.zip Rys_Szymon_2/ oswiadczenie.txt
	rm -r Rys_Szymon_2

.PHONY: lab03_zip
lab03_zip:
	mkdir Rys_Szymon_3
	mkdir -p archives
	cp -r lab03/ Rys_Szymon_3/
	zip -r archives/Rys_Szymon_3.zip Rys_Szymon_3/ oswiadczenie.txt
	rm -r Rys_Szymon_3

.PHONY: lab04_05_zip
lab04_zip:
	mkdir Rys_Szymon_4_5
	mkdir -p archives
	cp -r lab04-05/ Rys_Szymon_4_5/
	zip -r archives/Rys_Szymon_4_5.zip Rys_Szymon_4_5/ oswiadczenie.txt
	rm -r Rys_Szymon_4_5

.PHONY: lab06_zip
lab06_zip:
	mkdir Rys_Szymon_6
	mkdir -p archives
	cp -r lab06/ Rys_Szymon_6/
	zip -r archives/Rys_Szymon_6.zip Rys_Szymon_6/ oswiadczenie.txt
	rm -r Rys_Szymon_6

.PHONY: lab07_zip
lab07_zip:
	mkdir Rys_Szymon_7
	mkdir -p archives
	cp -r lab07/ Rys_Szymon_7/
	zip -r archives/Rys_Szymon_7.zip Rys_Szymon_7/ oswiadczenie.txt
	rm -r Rys_Szymon_7
