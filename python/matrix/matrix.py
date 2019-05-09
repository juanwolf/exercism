class Matrix(object):
    def __init__(self, matrix_string):
        self.rows = [
            [int(column_item) for column_item in row.split()]
            for row in matrix_string.splitlines()
        ]

    def row(self, index):
        return self.rows[index - 1]

    def column(self, index):
        return  [row[index - 1] for row in self.rows]
