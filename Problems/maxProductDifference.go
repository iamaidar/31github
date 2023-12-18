package problems

m1, m2, s1, s2 := 0, 0, 99999999999999, 99999999999999

 for i := 0; i < len(nums); i++ {
  if nums[i] > m1 {
   m2 = m1
   m1 = nums[i]
  } else if nums[i] > m2 {
   m2 = nums[i]
  }

  if nums[i] < s1 {
   s2 = s1
   s1 = nums[i]
  } else if nums[i] < s2 {
   s2 = nums[i]
  }
 }

    return m1*m2 - s1*s2 
}
