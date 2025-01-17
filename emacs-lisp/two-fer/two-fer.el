;;; two-fer.el --- Two-fer Exercise (exercism)

;;; Commentary:

;;; Code:
;;;
(defun two-fer (&optional name)
  "Return the two-fer for this name.
Example:
  (two-fer NAME) => One for NAME, one for me.
  (two-fer) => One for you, one for me."
  (unless name
    (setq name "you"))
  (message "One for %s, one for me." name))

(provide 'two-fer)
;;; two-fer.el ends here
