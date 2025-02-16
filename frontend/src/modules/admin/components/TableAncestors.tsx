import { ReactComponent as SendPublicIcon } from '@/assets/svg/public.svg';
import { Button, Chip, getKeyValue, Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@heroui/react';
import { Link } from '@tanstack/react-router';

const COLUMNS = [
  {
    name: 'NAME',
    uid: 'name',
  },
  {
    name: 'SURNAME',
    uid: 'surname',
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

const ROWS = [
  { name: 'Иван', surname: 'Иванов', status: 'На публикацию', actions: 'Упобликовать', view: { id: '1', lat: 1, lon: 1 } },
  { name: 'Петр', surname: 'Петров', status: 'На публикацию', actions: 'Упобликовать', view: { id: '2', lat: 2, lon: 2 } },
  { name: 'Сидор', surname: 'Сидоров', status: 'На публикацию', actions: 'Упобликовать', view: { id: '3', lat: 3, lon: 3 } },
  { name: 'Василий', surname: 'Васильев', status: 'На публикацию', actions: 'Упобликовать', view: { id: '4', lat: 4, lon: 4 } },
];

export const TableAncestors = () => {
  // const { data, isSuccess, isLoading, isError, error } = useOrder();
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
          <TableRow key={item.name}>
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
