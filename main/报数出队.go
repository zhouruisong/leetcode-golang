package main

/*
题目：13个人围坐一圈报数（123的报数），凡是报到3的人出队，问最后剩下的人的序号

JAVA实现
 * Created by Kameleon on 14-5-29.
 * 30个人围坐在一起报数，报到3的人出圈

public class Count {

	public static void main(String[] args){
		boolean[] arr = new boolean[13];//用于记录哪些人已经出队了
		for(int i = 0;i < arr.length;i++){
			arr[i] = true;
		}

		int leftCount = 13;
		int countNum = 0;
		int index = 0;//记录报数的人的序号，会超过13，所以之后的操作需要求余

		while(leftCount > 1){
			if(arr[index % 13] == true){
				countNum++;
				if(countNum == 3){
					countNum = 0;
					arr[index % 13] = false;
					leftCount--;
				}
			}
			index++;
		}

		//输出剩余的序号
		for(int i = 0;i < arr.length;i++){
			if(arr[i] == true){
				System.out.println(i);
			}
		}
	}
}

 */