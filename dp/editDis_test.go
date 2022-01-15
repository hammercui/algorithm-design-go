package dp

import "testing"

func Test_minDistance(t *testing.T) {
	//最小路径
	a := "horse"//ros
	b := "ros"
	t.Log("a",a)
	t.Log("b",b)

	ans,trailsMap := minDistance(a,b)

	trailsToStep(a,b,ans,trailsMap)
	t.Log("ans",ans)
	//{2,0,0}rorse
	//{1,2,2}rose
	//{1,4,2}ros
}
