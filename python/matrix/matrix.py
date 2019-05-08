class Matrix(object):
    def __init__(self, matrix_string):
        self.rows = []

        for row in matrix_string.splitlines():
            formatted_column = [
                int(column_item)
                for column_item in row.split()
            ]
            self.rows.append(formatted_column)

    def row(self, index):
        return self.rows[index - 1]

    def column(self, index):
        column = []
        for row in self.rows:
            column.append(row[index - 1])

        return column
