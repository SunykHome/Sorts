闲来无事复习了一下排序算法
使用golang语言编写
![image](https://user-images.githubusercontent.com/105403999/169494953-bc1212d8-11ec-468a-b3bd-f90cd666b2b7.png)
复杂度和稳定性一览
![1653037677(1)](https://user-images.githubusercontent.com/105403999/169495127-91b977cb-dc77-451d-bf3d-8c73f354e6f8.png)
冒泡排序
算法描述
对于要排序的数组，从第一位开始从前往后比较相邻两个数字，若前者大，则交换两数字位置，然后比较位向右移动一位。也就是比较arr[0]和arr[1]，若arr[0] > arr[1]，交换arr[0]和arr[1]。接着比较位移动一位，比较arr[1]和arr[2]，直到比较到arr[n - 2]和arr[n - 1] (n = arr.length)。第1轮从前到后的比较将使得最大的数字 冒泡 到最后，此时可以说一个数字已经被排序。每一轮的比较将使得当前未排序数字中的最大者被排序，未排序数字总数减1。第arr.length - 1轮结束后排序完成。
如下动图展示了{4,6,2,1,7,9,5,8,3}的冒泡排序过程（未应用优化）。
![image](https://user-images.githubusercontent.com/105403999/169495293-5288e31b-a598-4af9-861b-d81c8462bbf6.png)
稳定性：稳定。
排序 稳定性 指的是相等元素在原输入数组中的相对位置，在排序后不变，否则排序不稳定。稳定性对于纯数字数组来说意义不大，但对于包含多个属性的类数组来说，用于比较的属性之外其他属性不同时，保持排序的稳定性就很重要了。
冒泡排序始终只交换相邻元素，比较对象大小相等时不交换，相对位置不变，故稳定。
