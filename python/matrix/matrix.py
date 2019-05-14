class Matrix(object):
    """
    The Matrix represents a matrix containing a list of rows.
    :param matrix_string: The string representing the matrix.
    :type matrix_string: str
    :ivar rows: Contains the rows of the matrix
    :vartype rows: list(list(int))
    """
    def __init__(self, matrix_string):
        self.rows = [
            [int(column_item) for column_item in row.split()]
            for row in matrix_string.splitlines()
        ]

    def row(self, index):
        """
        Returns the row for a specific index.
        :param index: Index of the row
        """
        return self.rows[index - 1]

    def column(self, index):
        """
        Returns the column for a specific index.
        :param index: index of the column
        """
        return  [row[index - 1] for row in self.rows]
