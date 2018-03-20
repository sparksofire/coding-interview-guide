## 方案一 (Page 3)
```
=== RUN   TestStack
--- PASS: TestStack (0.00s)
	stack_test.go:11: 具有getMin 功能的栈（方案一）表组测试
	stack_test.go:30: 正在进行第1组测试, 共3组
	stack_test.go:33: 生成一组随机数，并依次插入栈中，随机数个数为 30，随机数最大值为 100
	stack_test.go:41: 随机数插入顺序为：[70 10 41 28 84 91 5 97 64 13 39 74 29 66 90 24 10 72 12 74 98 2 53 96 98 34 9 62 43 83]
	stack_test.go:43: 随机数从小到大排列为：[2 5 9 10 10 12 13 24 28 29 34 39 41 43 53 62 64 66 70 72 74 74 83 84 90 91 96 97 98 98]
	stack_test.go:52: 栈GetMin()返回值 2 与随机数最小值 2匹配
	stack_test.go:30: 正在进行第2组测试, 共3组
	stack_test.go:33: 生成一组随机数，并依次插入栈中，随机数个数为 50，随机数最大值为 200
	stack_test.go:41: 随机数插入顺序为：[178 129 143 48 106 9 26 141 36 58 178 67 162 123 183 99 25 60 182 2 98 106 167 13 59 18 117 46 97 158 54 12 142 84 48 61 12 122 55 173 128 176 143 173 56 98 166 29 108 166]
	stack_test.go:43: 随机数从小到大排列为：[2 9 12 12 13 18 25 26 29 36 46 48 48 54 55 56 58 59 60 61 67 84 97 98 98 99 106 106 108 117 122 123 128 129 141 142 143 143 158 162 166 166 167 173 173 176 178 178 182 183]
	stack_test.go:52: 栈GetMin()返回值 2 与随机数最小值 2匹配
	stack_test.go:30: 正在进行第3组测试, 共3组
	stack_test.go:33: 生成一组随机数，并依次插入栈中，随机数个数为 500，随机数最大值为 1000
	stack_test.go:41: 随机数插入顺序为：[698 272 967 289 302 965 494 115 9 418 720 476 282 420 279 842 490 822 95 666 667 523 249 114 526 660 798 658 975 236 860 991 143 724 705 434 859 416 670 325 443 743 854 247 737 310 678 252 201 6 642 616 584 218 251 932 598 57 27 814 223 369 665 411 490 240 597 631 813 614 217 577 233 848 182 410 36 313 408 117 967 930 821 836 993 425 382 65 258 969 840 883 517 213 700 431 325 222 851 895 248 130 706 446 75 480 285 752 845 73 81 954 360 59 606 872 207 376 975 213 167 894 792 703 905 687 855 738 520 970 237 403 840 141 408 57 781 569 927 944 363 506 142 121 305 158 922 753 515 308 940 339 905 402 974 426 363 521 600 183 293 879 436 437 67 227 653 215 6 703 818 205 741 330 91 91 531 267 63 670 669 945 368 63 163 485 853 213 936 640 680 229 379 510 154 70 710 779 923 719 296 309 293 700 548 891 659 644 32 799 831 223 352 274 375 743 479 230 630 697 619 26 32 403 680 253 247 697 346 457 371 715 225 864 899 521 50 384 632 292 734 768 819 537 438 590 599 217 342 563 198 825 793 60 288 529 852 373 222 823 387 300 659 252 679 143 200 841 185 677 783 776 385 528 87 866 525 323 823 611 618 57 590 937 412 520 690 698 80 516 736 734 603 756 937 412 528 720 936 383 43 139 168 916 552 545 77 236 86 324 938 591 642 964 246 320 164 555 192 323 291 892 150 675 290 807 426 500 447 783 42 761 321 175 882 785 971 307 589 144 822 975 396 409 396 614 331 966 822 758 811 425 147 758 961 854 651 239 124 155 438 472 513 371 740 552 398 611 729 982 144 840 111 805 370 63 409 424 71 247 926 903 213 474 969 271 508 541 835 780 666 775 39 64 831 495 385 845 468 697 320 351 188 485 87 954 672 664 174 362 126 141 30 807 179 23 156 250 109 467 700 623 230 687 849 350 857 820 837 410 712 346 207 690 376 179 961 813 951 471 598 299 348 32 23 477 879 599 187 387 192 783 877 663 357 518 966 293 302 897 426 769 771 266 860 534 620 113 956 916 851 926 691 698 775 190 672 620 313 900 941 327 108 635 697 481 454 312 956 846 957 128 241 371 990 127 586 38 747 175]
	stack_test.go:43: 随机数从小到大排列为：[6 6 9 23 23 26 27 30 32 32 32 36 38 39 42 43 50 57 57 57 59 60 63 63 63 64 65 67 70 71 73 75 77 80 81 86 87 87 91 91 95 108 109 111 113 114 115 117 121 124 126 127 128 130 139 141 141 142 143 143 144 144 147 150 154 155 156 158 163 164 167 168 174 175 175 179 179 182 183 185 187 188 190 192 192 198 200 201 205 207 207 213 213 213 213 215 217 217 218 222 222 223 223 225 227 229 230 230 233 236 236 237 239 240 241 246 247 247 247 248 249 250 251 252 252 253 258 266 267 271 272 274 279 282 285 288 289 290 291 292 293 293 293 296 299 300 302 302 305 307 308 309 310 312 313 313 320 320 321 323 323 324 325 325 327 330 331 339 342 346 346 348 350 351 352 357 360 362 363 363 368 369 370 371 371 371 373 375 376 376 379 382 383 384 385 385 387 387 396 396 398 402 403 403 408 408 409 409 410 410 411 412 412 416 418 420 424 425 425 426 426 426 431 434 436 437 438 438 443 446 447 454 457 467 468 471 472 474 476 477 479 480 481 485 485 490 490 494 495 500 506 508 510 513 515 516 517 518 520 520 521 521 523 525 526 528 528 529 531 534 537 541 545 548 552 552 555 563 569 577 584 586 589 590 590 591 597 598 598 599 599 600 603 606 611 611 614 614 616 618 619 620 620 623 630 631 632 635 640 642 642 644 651 653 658 659 659 660 663 664 665 666 666 667 669 670 670 672 672 675 677 678 679 680 680 687 687 690 690 691 697 697 697 697 698 698 698 700 700 700 703 703 705 706 710 712 715 719 720 720 724 729 734 734 736 737 738 740 741 743 743 747 752 753 756 758 758 761 768 769 771 775 775 776 779 780 781 783 783 783 785 792 793 798 799 805 807 807 811 813 813 814 818 819 820 821 822 822 822 823 823 825 831 831 835 836 837 840 840 840 841 842 845 845 846 848 849 851 851 852 853 854 854 855 857 859 860 860 864 866 872 877 879 879 882 883 891 892 894 895 897 899 900 903 905 905 916 916 922 923 926 926 927 930 932 936 936 937 937 938 940 941 944 945 951 954 954 956 956 957 961 961 964 965 966 966 967 967 969 969 970 971 974 975 975 975 982 990 991 993]
	stack_test.go:52: 栈GetMin()返回值 6 与随机数最小值 6匹配
PASS
ok  	github.com/shawpo/codingInterviewGuide/chapter1/getMinStack/stack1	0.008s
```