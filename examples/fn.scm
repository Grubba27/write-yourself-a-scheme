(fn
    fold (list fn acc)
    (if (< (len list) 1)
        acc
    (fold (tail list) (fn (head list) acc))))
(fn sum (a b) (+ a b))
(fold 1 sum 0)