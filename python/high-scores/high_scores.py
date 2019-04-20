class HighScores(object):
    def __init__(self, scores):
        self.scores = scores

    def personal_top_three(self):
        return self._get_best_scores(3)

    def latest(self):
        return self.scores[-1]

    def personal_best(self):
        return self._get_best_scores(1)[0]

    def _get_best_scores(self, n):
        """_get_best_scores returns the n best scores from this instance of HighScores"""
        return sorted(self.scores, reverse=True)[0:n]
