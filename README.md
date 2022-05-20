闲来无事复习了一下排序算法<br />使用golang语言编写<br />![](https://cdn.nlark.com/yuque/0/2022/png/26902243/1653038029548-ce8a65fb-1d71-4473-8b11-8ae60b919de4.png#clientId=ub2f572c0-7ff7-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=442&id=ue05a91d6&originHeight=888&originWidth=1236&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u6eb6b59b-bbce-4870-8d99-6ab307b11e1&title=&width=615)<br />**复杂度和稳定性一览**<br />![1653038073(1).jpg](https://cdn.nlark.com/yuque/0/2022/jpeg/26902243/1653038082089-36af2bf6-6405-43cb-a6c3-0c96fa3cf118.jpeg#clientId=ub2f572c0-7ff7-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=572&id=u95bdde56&name=1653038073%281%29.jpg&originHeight=776&originWidth=877&originalType=binary&ratio=1&rotation=0&showTitle=false&size=67306&status=done&style=none&taskId=u42a4942a-9904-429c-8ab5-18a42849883&title=&width=646.6000366210938)

> **冒泡排序**<br />算法描述<br />对于要排序的数组，从第一位开始从前往后比较相邻两个数字，若前者大，则交换两数字位置，然后比较位向右移动一位。也就是比较arr[0]和arr[1]，若arr[0] > arr[1]，交换arr[0]和arr[1]。接着比较位移动一位，比较arr[1]和arr[2]，直到比较到arr[n - 2]和arr[n - 1] (n = arr.length)。第1轮从前到后的比较将使得最大的数字 冒泡 到最后，此时可以说一个数字已经被排序。每一轮的比较将使得当前未排序数字中的最大者被排序，未排序数字总数减1。第arr.length - 1轮结束后排序完成。<br />如下动图展示了{4,6,2,1,7,9,5,8,3}的冒泡排序过程（未应用优化）。<br />![](https://cdn.nlark.com/yuque/0/2022/gif/26902243/1653038044860-412e9b50-9573-47f5-812e-14fc0c15bc2f.gif#clientId=ub2f572c0-7ff7-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u4b81b012&originHeight=189&originWidth=344&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u7ddbff71-4451-4873-9850-b1314ab153c&title=)<br />稳定性：稳定。<br />排序 稳定性 指的是相等元素在原输入数组中的相对位置，在排序后不变，否则排序不稳定。稳定性对于纯数字数组来说意义不大，但对于包含多个属性的类数组来说，用于比较的属性之外其他属性不同时，保持排序的稳定性就很重要了。<br />冒泡排序始终只交换相邻元素，比较对象大小相等时不交换，相对位置不变，故稳定。
> 
> **选择排序**
> 算法描述
> 对于要排序的数组，设置一个minIdx记录最小数字下标。先假设第1个数字最小，此时minIdx = 0，将arr[minIdx]与后续数字逐一比较，当遇到更小的数字时，使minIdx等于该数字下标，第1轮比较将找出此时数组中最小的数字。找到后将minIdx下标的数字与第1个数字交换，此时称一个数字已被排序。然后开始第2轮比较，令minIdx = 1，重复上述过程。每一轮的比较将使得当前未排序数字中的最小者被排序，未排序数字总数减1。第arr.length - 1轮结束后排序完成。
> 如下动图展示了{4,6,2,1,7,9,5,8,3}的选择排序过程（单元选择）。
> ![](https://cdn.nlark.com/yuque/0/2022/gif/26902243/1653038212608-0d6f90fc-849b-4085-a094-34a31f8edf18.gif#clientId=ub2f572c0-7ff7-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u4288549d&originHeight=189&originWidth=344&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=ucc540c3c-11d6-45a3-9624-27bcbb30730&title=)
> 微优化：在交换前判断minIdx是否有变化，若无变化则无需交换。当数组大致有序时，能够减少无效交换带来的开销。
> 稳定性：不稳定。
> 存在跨越交换。找到本轮次最小值之后，将其与本轮起始数字交换，此时若中间有与起始元素同值的元素，将打破稳定性
> 插入排序
> 算法描述
> 对于待排序数组，从第2个元素开始(称作插入对象元素)，比较它与之前的元素(称作比较对象元素)，当插入对象元素小于比较对象元素时，继续往前比较，直到不小于(≥)比较对象，此时将插入对象元素插入到该次比较对象元素之后。重复这个插入过程直到最后一个元素作为插入对象元素完成插入操作。
> 如下动图展示了{4,6,2,1,7,9,5,8,3}的简单插入排序过程。
> ![](https://cdn.nlark.com/yuque/0/2022/gif/26902243/1653038324377-00c46c6d-ff6d-49a6-bb09-821fcaa59747.gif#clientId=ub2f572c0-7ff7-4&crop=0&crop=0&crop=1&crop=1&from=paste&id=u4eae3341&originHeight=382&originWidth=344&originalType=url&ratio=1&rotation=0&showTitle=false&status=done&style=none&taskId=u419f9ed8-eb80-4697-9372-8c8c99e9366&title=)

> 稳定性：简单插入和折半插入(二分插入)排序是稳定的。


