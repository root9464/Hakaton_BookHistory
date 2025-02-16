import { Button, Modal, ModalBody, ModalContent, ModalFooter, ModalHeader } from '@heroui/react';
import { Point } from 'framer-motion';
import { useAncestor } from '../hook/useAncestor';

type AncestorModalProps = {
  isOpen: boolean;
  onOpen: () => void;
  onOpenChange: () => void;
  coordinates: Point & { id: string };
};

export const AncestorModal = ({ isOpen, onOpen, onOpenChange, coordinates }: AncestorModalProps) => {
  const { data } = useAncestor(coordinates.id);

  return (
    <>
      <button onClick={onOpen} className='data?s-center flex h-fit w-fit flex-row justify-center bg-transparent text-white'>
        Create
      </button>
      <Modal isOpen={isOpen} onOpenChange={onOpenChange} size='full'>
        <ModalContent className='z-[2]'>
          {(onClose) => (
            <>
              <ModalHeader className='flex flex-col gap-1'>Информация о герое</ModalHeader>
              <ModalBody>
                <p>{data?.fio}</p>
                <p>{data?.date_of_birth}</p>
                <p>{data?.place_of_birth}</p>
                <p>{data?.name_of_millitary_commissariat}</p>
                <p>{data?.millitary_rank}</p>
                <p>{data?.date_of_death}</p>
                <p>{data?.burial_place}</p>
                <p>{data?.biographical_facts}</p>
              </ModalBody>
              <ModalFooter>
                <Button color='danger' variant='light' onPress={onClose}>
                  Close
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </>
  );
};
