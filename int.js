let arr = [0, 4, 2, 1, 3];
var numb = 0;
for (i = 0; i < arr.length; i++) {
  for (k = 0; k < i; k++) {
    if (arr[i] > arr[k]) {
      arr[k] = arr[i];
      arr[i] = arr[k];
    }
    console.log(arr[numb]);
    console.log(arr);
  }
  numb++;
}
