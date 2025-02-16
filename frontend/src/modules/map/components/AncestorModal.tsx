import { Button, Modal, ModalBody, ModalContent, ModalFooter, ModalHeader } from '@heroui/react';
import { ReactNode } from '@tanstack/react-router';
import { useAncestor } from '../hook/useAncestor';
import { Point } from '../Module';

type AncestorModalProps = {
  isOpen: boolean;
  onOpenChange: () => void;
  coordinates: Point & { id: string };
  children?: ReactNode;
};

export const AncestorModal = ({ isOpen, onOpenChange, coordinates, children }: AncestorModalProps) => {
  const { data } = useAncestor(coordinates.id);

  return (
    <>
      {children}
      <Modal isOpen={isOpen} onOpenChange={onOpenChange} size='full' className='z-[1]' inert>
        <ModalContent>
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
