#作业: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

#思路 sql.ErrNoRows作为标准库的sentinel 在dao这一层应该Wrap这个error，交给上层做处理。
