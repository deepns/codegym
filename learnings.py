# noting down some learnings along the way

def func_args_init():
    # the argument values is initialized with an empty list when interpreted.
    # subsequent calls to f() will use the same object assigned to values assigned during initialization
    def f(i, values = []):
        values.append(i)
        print(values)
        return values
    
    # will print [1], [1, 2], [1, 2, 3]
    f(1)
    f(2)
    f(3)

func_args_init()