package xsql2

import "bytes"

func (order *XSql2Order)Count()int64{
	if  len(order.tables) <= 0 {
		return 0
	}
	var sqlOrder bytes.Buffer
	sqlOrder.WriteString( "select count(*) from ")
	var pos,joinIndex int
	pos = -1
	if len(order.join)>0{
		pos=order.join[0].pos
	}
	for index, _ := range order.tables {
		sqlOrder.WriteString( order.tables[index].GetName())
		if pos!= -1 && index == pos && joinIndex < len(order.join){
			for {
				if order.join[index].pos == pos && joinIndex < len(order.join) {
					sqlOrder.WriteString(order.getJoinString(joinIndex))
					joinIndex++
				}else {
					break
				}
			}
		}
		if index != len(order.tables)-1 {
			sqlOrder.WriteString( " , ")
		}
	}
	sqlOrder.WriteString( order.getWhereString())
	sqlOrder.WriteString( order.getOrderString())
	n := order.executeForCount(sqlOrder.String())


	return n
}
