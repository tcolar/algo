# You run an e-commerce website and want to record the last N order ids in a log. Implement a data structure to accomplish this, with the following API:

#  * record(order_id): adds the order_id to the log
#  * get_last(i): gets the ith last element from the log. i is guaranteed to be smaller than or equal to N.
# You should be as efficient with time and space as possible.

class log:
    orders = []
    tail = 0

    def __init__(self, size):
        self.size = size
        self.orders = size * [None]

    def record(self, id):
        self.orders[self.tail] = id
        self.tail += 1
        if self.tail >= self.size:
            self.tail = 0

    def get_last(self, i):
        index = self.tail - i
        if index < 0:
            index = self.size + index
        print("{}  --".format(index))
        return self.orders[index]


l = log(6)
l.record(1)
l.record(2)
l.record(3)
l.record(4)
l.record(5)
l.record(6)
l.record(7)
l.record(8)
print(l.orders)
print(l.get_last(1))
print(l.get_last(5))
