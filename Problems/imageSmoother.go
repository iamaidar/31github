package problems

func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := range img {
		res[i] = make([]int, len(img[i]))
	}

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[0]); j++ {
			sum := 0
			count := 0
			for n := i - 1; n < i+2; n++ {
				for m := j - 1; m < j+2; m++ {
					if n >= 0 && m >= 0 && n < len(img) && m < len(img[0]) {
						count++
						sum += img[n][m]
					}
				}
			}
			res[i][j] = sum / count
		}
	}

	return res
}

/*
func imageSmoother(img [][]int) [][]int {
    var smooth [][]int = make([][]int, len(img))
    for i := 0; i<len(img); i++{
        smooth[i] = make([]int, len(img[0]))
    }
    l := len(img)    //length
    w := len(img[0])    //width
    for i := 0; i < l; i++ {
        for j := 0; j < w ; j++{
            sum := 0
            n := 0
            //above left
            if i > 0 && j > 0{
                n+=1
                sum += img[i-1][j-1]
            }
            //above
            if i > 0 {
                n+=1
                sum += img[i-1][j]
            }
            //above right
            if i > 0 && j < w-1{
                n+=1
                sum += img[i-1][j+1]
            }

            // left
            if j > 0{
                n+=1
                sum += img[i][j-1]
            }
            //self
            n+=1
            sum += img[i][j]
            // right
            if j < w-1{
                n+=1
                sum += img[i][j+1]
            }
            // below left
            if i < l-1 && j > 0{
                n+=1
                sum += img[i+1][j-1]
            }
            //below
            if i < l-1 {
                n+=1
                sum += img[i+1][j]
            }
            //below right
            if i < l-1 && j < w-1{
                n+=1
                sum += img[i+1][j+1]
            }
            //calculate avg
            smooth[i][j] = sum/n
        }
    }
    return smooth
}
*/
