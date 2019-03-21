//处理函数可以是系统提供的处理函数，如将字符串变大写或小写，也可以使用自定义函数。
//本例中的字符串处理函数的逻辑是使用一个自定义的函数实现一出指定go前缀的过程

//自定义的一出前缀的处理函数
func removePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

//此函数使用了strings.TrimPrefix()函数视线一出字符串的制定前缀。处理后，移除前缀的字符串结果将通过
//removePrefix()函数的返回值返回