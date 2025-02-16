import { ReactComponent as SendPublicIcon } from '@/assets/svg/public.svg';
import { Button, Chip, getKeyValue, Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@heroui/react';
import { Link } from '@tanstack/react-router';
import { useCombinedOrders } from '../hooks/useOrder';

const COLUMNS = [
  {
    name: 'FIO',
    uid: 'fio',
  },
  {
    name: 'STATUS',
    uid: 'status',
  },
  {
    name: 'ACTIONS',
    uid: 'actions',
  },
  {
    name: 'View',
    uid: 'view',
  },
];

export const TableAncestors = () => {
  const { data } = useCombinedOrders();
  const ROWS = data
    ? data.map((order) => {
        const [lon, lat] = order.geom.match(/\d+\.\d+/g)?.map(Number) || [0, 0];
        return {
          fio: order.fio,
          status: order.status === 'draft' ? 'На публикацию' : order.status,
          actions: 'Упобликовать',
          view: { id: order.id, lat, lon },
        };
      })
    : [];

  return (
    <Table
      aria-label='Example table with dynamic content'
      selectionBehavior='toggle'
      classNames={{
        base: 'w-full shadow-none',
        wrapper: 'shadow-none',
        table: 'shadow-none',
      }}
    >
      <TableHeader columns={COLUMNS}>
        {(column) => (
          <TableColumn key={column.uid} className='border-none bg-uiDeepGray text-white'>
            {column.name}
          </TableColumn>
        )}
      </TableHeader>
      <TableBody items={ROWS}>
        {(item) => (
          <TableRow key={item.fio}>
            {(columnKey) => (
              <TableCell className='w-max'>
                {columnKey === 'actions' && (
                  <Button className='flex items-center gap-2 text-white' color='success'>
                    <SendPublicIcon className='h-5 w-5 stroke-white' />
                    <span>{getKeyValue(item, columnKey)}</span>
                  </Button>
                )}
                {columnKey === 'status' && (
                  <Chip className='capitalize' color='success' size='sm' variant='flat'>
                    {getKeyValue(item, columnKey)}
                  </Chip>
                )}
                {columnKey === 'view' && (
                  <Button as={Link} to='/map' className='flex w-fit items-center gap-2 bg-uiDeepGray text-white' color='success'>
                    Просмотреть
                  </Button>
                )}
                {!['actions', 'status', 'view'].includes(columnKey.toString()) && <span>{getKeyValue(item, columnKey)}</span>}{' '}
              </TableCell>
            )}
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
};
