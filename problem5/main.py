def cons(a, b):
    def pair(f):
        return f(a, b)

    return pair


def car(pair):
    def pair_func(a, b): return a, b

    final_pair = pair(pair_func)
    return final_pair[0]


def cdr(pair):
    def pair_func(a, b): return a, b

    final_pair = pair(pair_func)
    return final_pair[1]


print(car(cons(1, 2)))
print(cdr(cons(1, 2)))
