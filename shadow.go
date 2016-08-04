

// No -Wshadow. This has caused problems many times for us. Introduction of new variables that shadow older variables can be problematic, particularly for functions that return multiple values. It is too easy to refactor something and clobber a pre-existing variable in some inner loop. -Wshadow is a
// much safer default. Perhaps some controlled de-scoping of a variable would be nice, or perhaps introducing an operater that allows persistence of a variable outside an if statement, like this:
if $resultToKeep, err := someOperation(foo); err != nil {
	// error
}
// resultToKeep remains in scope but err doesn't. The alternative is to have linear or cascaded if/else statements, which is ugly, and falls apart as soon as you need to add a loop or something. One nice heuristic that I would like to try is that small names (tmp, err, ok) are allowed to persist and
// clobber each other, but >3 char variables may not clobber. I'll bet that 'i', 'err' and 'ok' are the most common variable names in most codebases.
