
import React from 'react'

const Table = ({head,children}:{head:React.ReactNode,children: React.ReactNode }) => {
  return (
      <div className="relative overflow-x-auto rounded-sm">
    <table className="w-full text-sm text-left rtl:text-right text-gray-500 rounded-md"> 
        <thead className="text-xs text-gray-800 uppercase bg-gray-300">
                  {
            head
            }
        </thead>
        <tbody className='text-gray-800'>
              {children}
        </tbody>
    </table>
</div>
  )
}
const TableHead = function({children}:{children: React.ReactNode }) {
  return (
    <th scope="col" className="px-6 py-3">
        {children}
                </th>
  )
}
const TableRow = function ({ children }: { children: React.ReactNode }) { 
    return (
        <tr className= "border-b odd:bg-white even:bg-gray-100">

            {children}

            </tr>
    )
}

function TableData({ children }: { children: React.ReactNode }) {
    return (
<td className="px-6 py-4">
    {children}
                </td>
    )
}
Table.Head = TableHead
Table.Row = TableRow
Table.Data = TableData

export default Table