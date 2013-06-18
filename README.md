wipi
====

Wiring Wrapper in Go for the Raspberry Pi

First build and install WiringPi.

	cd WiringPi/wiringPi
	make
	sudo make install


	+----------+------+--------+-------+
	| wiringPi | GPIO | Name   | Value |
	+----------+------+--------+-------+
	|      0   |  17  | GPIO 0 | Low   |
	|      1   |  18  | GPIO 1 | High  |
	|      2   |  21  | GPIO 2 | Low   |
	|      3   |  22  | GPIO 3 | Low   |
	|      4   |  23  | GPIO 4 | Low   |
	|      5   |  24  | GPIO 5 | Low   |
	|      6   |  25  | GPIO 6 | Low   |
	|      7   |   4  | GPIO 7 | High  |
	|      8   |   0  | SDA    | High  |
	|      9   |   1  | SCL    | High  |
	|     10   |   8  | CE0    | Low   |
	|     11   |   7  | CE1    | Low   |
	|     12   |  10  | MOSI   | Low   |
	|     13   |   9  | MISO   | Low   |
	|     14   |  11  | SCLK   | Low   |
	|     15   |  14  | TxD    | High  |
	|     16   |  15  | RxD    | High  |
	+----------+------+--------+-------+
