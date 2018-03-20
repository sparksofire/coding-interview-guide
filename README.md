# Coding Interview Guide
> Golang 重写左程云著《程序员代码面试指南》一书中的示例代码(原书为JAVA)

## 代码结构
- 根目录下的chapterX表示第X章
- chapterX下的目录名表示示例代码的题目名称
- 题目下的目录表示该题目的不同解决方案（若存在）
- .go为具体代码，其中xxx.go格式的文件为实现代码，而xxx_test.go则表示对应的单元测试。README.md为说明文件

## 单元测试
如代码结构所述，每个题目都有自己的单元测试，进入所在目录，执行go test -v即可运行单元测试。

例如chapter1/getMinStack目录下的getMin_test.go对该题目的两种解决方案进行了测试，以下为执行go test -v的输出结果
```
=== RUN   TestStack
--- PASS: TestStack (0.00s)
	getMin_test.go:13: 具有getMin 功能的栈表组测试
	getMin_test.go:32: 正在进行第1组测试, 共3组
	getMin_test.go:36: 生成一组随机数，并依次插入栈中，随机数个数为 30，随机数最大值为 100
	getMin_test.go:45: 随机数插入顺序为：[74 30 77 81 31 56 31 56 96 93 79 65 14 70 61 39 86 68 89 59 25 99 9 3 69 24 6 52 4 32]
	getMin_test.go:47: 随机数从小到大排列为：[3 4 6 9 14 24 25 30 31 31 32 39 52 56 56 59 61 65 68 69 70 74 77 79 81 86 89 93 96 99]
	getMin_test.go:57: 方案一栈GetMin()返回值 3 与随机数最小值 3匹配
	getMin_test.go:67: 方案二栈GetMin()返回值 3 与随机数最小值 3匹配
	getMin_test.go:32: 正在进行第2组测试, 共3组
	getMin_test.go:36: 生成一组随机数，并依次插入栈中，随机数个数为 50，随机数最大值为 200
	getMin_test.go:45: 随机数插入顺序为：[50 63 91 72 159 172 42 32 195 198 15 3 152 103 22 81 138 26 38 64 133 61 44 195 88 22 29 90 20 149 9 126 140 22 70 55 164 94 17 105 27 146 191 100 79 167 39 129 88 100]
	getMin_test.go:47: 随机数从小到大排列为：[3 9 15 17 20 22 22 22 26 27 29 32 38 39 42 44 50 55 61 63 64 70 72 79 81 88 88 90 91 94 100 100 103 105 126 129 133 138 140 146 149 152 159 164 167 172 191 195 195 198]
	getMin_test.go:57: 方案一栈GetMin()返回值 3 与随机数最小值 3匹配
	getMin_test.go:67: 方案二栈GetMin()返回值 3 与随机数最小值 3匹配
	getMin_test.go:32: 正在进行第3组测试, 共3组
	getMin_test.go:36: 生成一组随机数，并依次插入栈中，随机数个数为 500，随机数最大值为 1000
	getMin_test.go:45: 随机数插入顺序为：[962 341 765 202 404 19 768 679 204 28 759 823 175 199 50 645 910 640 53 405 803 75 358 44 220 532 521 298 192 918 249 595 18 890 614 497 155 586 786 110 143 198 641 62 149 369 25 785 633 232 796 885 199 601 670 580 634 959 922 727 99 74 338 595 456 61 605 701 14 623 866 326 725 4 747 381 574 365 755 384 801 215 956 248 432 532 53 857 938 754 356 963 220 492 418 980 862 748 843 197 691 116 82 256 670 563 459 488 406 968 338 544 654 292 801 86 15 77 781 932 902 310 865 769 445 546 394 961 493 38 339 94 490 735 113 559 824 200 524 697 633 334 385 659 28 941 602 453 265 894 437 464 369 671 900 76 854 122 302 179 282 938 690 880 42 361 287 422 984 319 728 567 492 627 74 280 298 795 91 806 183 359 324 528 187 846 745 558 418 929 653 914 33 949 293 863 710 791 850 442 188 345 740 944 758 810 125 952 92 695 611 139 566 843 845 520 711 82 354 61 566 839 635 353 600 65 779 646 160 100 330 294 895 664 795 150 481 571 358 246 588 207 703 156 258 106 435 897 439 568 865 577 656 605 597 675 837 655 345 321 505 46 1 491 670 425 492 934 748 849 960 564 892 620 162 703 900 855 664 706 860 917 406 892 698 545 449 688 443 27 519 546 865 478 25 697 557 500 473 875 201 760 452 62 615 918 232 701 779 796 406 424 834 876 617 662 846 733 892 781 974 352 533 557 65 739 543 195 177 570 121 573 943 450 40 877 435 900 715 272 402 386 436 564 705 253 464 370 284 56 295 274 338 900 728 311 788 412 960 832 573 925 231 905 469 492 952 765 560 37 64 826 57 231 311 737 812 642 581 394 151 81 53 300 523 18 773 949 685 961 362 591 935 205 995 775 997 49 28 665 612 833 303 698 134 413 147 281 770 329 969 494 34 481 672 332 526 675 171 403 64 685 107 84 250 313 541 459 415 361 558 947 65 275 576 400 812 539 230 80 879 426 562 413 195 666 497 108 282 521 220 423 138 218 465 260 280 822 970 849 192 38 465 679 723 672 132 452 486 577 805 242 734 163 148 365 593 80 204 74 967 432 14 639 378 672 678 455 652 978 554 261 944 138 822 781 721 881 167 803]
	getMin_test.go:47: 随机数从小到大排列为：[1 4 14 14 15 18 18 19 25 25 27 28 28 28 33 34 37 38 38 40 42 44 46 49 50 53 53 53 56 57 61 61 62 62 64 64 65 65 65 74 74 74 75 76 77 80 80 81 82 82 84 86 91 92 94 99 100 106 107 108 110 113 116 121 122 125 132 134 138 138 139 143 147 148 149 150 151 155 156 160 162 163 167 171 175 177 179 183 187 188 192 192 195 195 197 198 199 199 200 201 202 204 204 205 207 215 218 220 220 220 230 231 231 232 232 242 246 248 249 250 253 256 258 260 261 265 272 274 275 280 280 281 282 282 284 287 292 293 294 295 298 298 300 302 303 310 311 311 313 319 321 324 326 329 330 332 334 338 338 338 339 341 345 345 352 353 354 356 358 358 359 361 361 362 365 365 369 369 370 378 381 384 385 386 394 394 400 402 403 404 405 406 406 406 412 413 413 415 418 418 422 423 424 425 426 432 432 435 435 436 437 439 442 443 445 449 450 452 452 453 455 456 459 459 464 464 465 465 469 473 478 481 481 486 488 490 491 492 492 492 492 493 494 497 497 500 505 519 520 521 521 523 524 526 528 532 532 533 539 541 543 544 545 546 546 554 557 557 558 558 559 560 562 563 564 564 566 566 567 568 570 571 573 573 574 576 577 577 580 581 586 588 591 593 595 595 597 600 601 602 605 605 611 612 614 615 617 620 623 627 633 633 634 635 639 640 641 642 645 646 652 653 654 655 656 659 662 664 664 665 666 670 670 670 671 672 672 672 675 675 678 679 679 685 685 688 690 691 695 697 697 698 698 701 701 703 703 705 706 710 711 715 721 723 725 727 728 728 733 734 735 737 739 740 745 747 748 748 754 755 758 759 760 765 765 768 769 770 773 775 779 779 781 781 781 785 786 788 791 795 795 796 796 801 801 803 803 805 806 810 812 812 822 822 823 824 826 832 833 834 837 839 843 843 845 846 846 849 849 850 854 855 857 860 862 863 865 865 865 866 875 876 877 879 880 881 885 890 892 892 892 894 895 897 900 900 900 900 902 905 910 914 917 918 918 922 925 929 932 934 935 938 938 941 943 944 944 947 949 949 952 952 956 959 960 960 961 961 962 963 967 968 969 970 974 978 980 984 995 997]
	getMin_test.go:57: 方案一栈GetMin()返回值 1 与随机数最小值 1匹配
	getMin_test.go:67: 方案二栈GetMin()返回值 1 与随机数最小值 1匹配
PASS
ok  	github.com/shawpo/coding-interview-guide/chapter1/getMinStack	0.008s

```
## 数据结构
为了方便，数据结构目前采用开源项目 github.com/emirpasic/gods，后面会进一步优化
