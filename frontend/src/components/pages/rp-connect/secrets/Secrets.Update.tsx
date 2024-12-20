import {Button, ButtonGroup, createStandaloneToast, Flex, FormField, Input, PasswordInput} from '@redpanda-data/ui';
import {PageComponent, PageInitHelper} from '../../Page';
import {observer} from 'mobx-react';
import {appGlobal} from '../../../../state/appGlobal';
import {pipelinesApi, rpcnSecretManagerApi} from '../../../../state/backendApi';
import PageContent from '../../../misc/PageContent';
import {action, makeObservable, observable} from 'mobx';
import {DefaultSkeleton} from '../../../../utils/tsxUtils';
import {formatPipelineError} from '../errors';
import {Scope, UpdateSecretRequest} from '../../../../protogen/redpanda/api/dataplane/v1alpha2/secret_pb';

const {ToastContainer, toast} = createStandaloneToast();

const returnSecretTab = '/connect-clusters?defaultTab=redpanda-connect-secret'

@observer
class RpConnectSecretUpdate extends PageComponent<{ secretId: string }> {
    @observable secret = '';
    @observable isUpdating = false;

    constructor(p: any) {
        super(p);
        makeObservable(this, undefined, {autoBind: true});
    }

    initPage(p: PageInitHelper) {
        p.title = 'Update Secret';
        p.addBreadcrumb('Redpanda Connect Secret Manager', '/rp-connect/secrets/update');
        p.addBreadcrumb('Update Secret', '');

        this.refreshData(true);
        appGlobal.onRefresh = () => this.refreshData(true);
    }

    refreshData(_force: boolean) {
        rpcnSecretManagerApi.refreshSecrets(_force);
    }

    cancel() {
        this.secret = '';
        appGlobal.history.push(returnSecretTab);
    }

    async updateSecret() {
        this.isUpdating = true;

        //create function given string return base64 encoded Uint8Array
        function base64Encode(str: string): Uint8Array {
            const encodedString = btoa(str);
            const charList = encodedString.split('').map(char => char.charCodeAt(0));
            return new Uint8Array(charList);
        }


        rpcnSecretManagerApi.update(this.props.secretId, new UpdateSecretRequest({
            id: this.props.secretId,
            secretData: base64Encode(this.secret),
            scopes: [Scope.REDPANDA_CONNECT]
        }))
            .then(async () => {
                toast({
                    status: 'success', duration: 4000, isClosable: false,
                    title: 'Secret updated'
                });
                await pipelinesApi.refreshPipelines(true);
                appGlobal.history.push(returnSecretTab);
            })
            .catch(err => {
                toast({
                    status: 'error', duration: null, isClosable: true,
                    title: 'Failed to update secret',
                    description: formatPipelineError(err),
                })
            })
            .finally(() => {
                this.isUpdating = false;
            });
    }

    render() {
        if (!rpcnSecretManagerApi.secrets) return DefaultSkeleton;

        const isSecretEmpty = this.secret.trim().length == 0;

        return (
            <PageContent>
                <ToastContainer/>
                <Flex flexDirection="column" gap={5}>
                    <FormField label="Secret name">
                        <Flex alignItems="center" gap="2">
                            <Input
                                placeholder="Enter a secret name..."
                                data-testid="secretId"
                                pattern="^[A-Z][A-Z0-9_]*$"
                                isRequired
                                disabled={true}
                                value={this.props.secretId}
                                width={500}
                            />
                        </Flex>
                    </FormField>

                    <FormField label="Secret value">
                        <Flex alignItems="center" gap="2" width={500}>
                            <PasswordInput
                                placeholder="Enter a new secret value..."
                                data-testid="secretValue"
                                isRequired
                                value={this.secret}
                                onChange={x => this.secret = x.target.value}
                                width={500}
                                type="password"
                                isDisabled={this.isUpdating}
                            />
                        </Flex>
                    </FormField>

                    <ButtonGroup>
                        <Button isLoading={this.isUpdating} isDisabled={isSecretEmpty}
                                onClick={action(() => this.updateSecret())}
                        >
                            Update Secret
                        </Button>
                        <Button variant="link" disabled={this.isUpdating} onClick={action(() => this.cancel())}>
                            Cancel
                        </Button>
                    </ButtonGroup>
                </Flex>
            </PageContent>
        );

    }
}

export default RpConnectSecretUpdate;
