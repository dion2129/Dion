package main
import "fmt"
func main () {
	studentlist := printStudent("a","b","c","d")
	result := sumNumber(studentlist...)
	fmt.Println(result)
}
func sumNumber(number ...int) int{
	total:= 0

	for _, value := range number{
		total += value
		}
	return total
}
func printStudent(names ...string) []string {
	var result []string
	for _, value := range names{
		result = append(result, value)
		}
	return result
}